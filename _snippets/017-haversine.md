---
title: Haversine
---

```php
<?php

declare(strict_types=1);

namespace Task_2\Exceptions;

use Exception;
use Throwable;

class OutOfRangeLatitude extends Exception
{
    private const MESSAGE = 'Latitude must be with the -90°,90° range';

    public function __construct(string $message = "", int $code = 0, ?Throwable $previous = null)
    {
        $message = self::MESSAGE . ': ' . $message;
        parent::__construct($message, $code, $previous);
    }
}
```

```php
<?php

declare(strict_types=1);

namespace Task_2\Exceptions;

use Exception;
use Throwable;

class OutOfRangeLongitude extends Exception
{
    private const MESSAGE = 'Longitude must be with the -180°,180° range';

    public function __construct(string $message = "", int $code = 0, ?Throwable $previous = null)
    {
        $message = self::MESSAGE . ': ' . $message;
        parent::__construct($message, $code, $previous);
    }
}
```

```php
<?php

declare(strict_types=1);

namespace Task_2;

use Task_2\Exceptions\OutOfRangeLatitude;
use Task_2\Exceptions\OutOfRangeLongitude;

class Haversine
{
    private const EARTH_RADIUS = 6371000;
    private const COVERAGE_LIMIT = 10000;

    public function __construct(
        private float $lat1,
        private float $lng1,
        private float $lat2,
        private float $lng2
    ) {
        $this->validatesInput();
        $this->convertsDegreesToRadiants();
    }

    private function validatesInput(): void {
        if ($this->lat1 < -90 || $this->lat1 > 90) {
            throw new OutOfRangeLatitude((string)$this->lat1);
        }
        if ($this->lat2 < -90 || $this->lat2 > 90) {
            throw new OutOfRangeLatitude((string)$this->lat2);
        }

        if ($this->lng1 < -180 || $this->lng1 > 180) {
            throw new OutOfRangeLongitude((string)$this->lng1);
        }
        if ($this->lng2 < -180 || $this->lng2 > 180) {
            throw new OutOfRangeLongitude((string)$this->lng2);
        }
    }

    private function convertsDegreesToRadiants(): void {
        $this->lat1 = deg2rad($this->lat1);
        $this->lat2 = deg2rad($this->lat2);
        $this->lng1 = deg2rad($this->lng1);
        $this->lng2 = deg2rad($this->lng2);
    }

    /**
     * Calculate the distance between the given coordinates using the Haversine formula.
     *
     * The formula is shamelessly stolen from wikipedia (https://en.wikipedia.org/wiki/Haversine_formula).
     * The resulting value might be right or not.
     * At first glance, nothing indicates something deeply wrong.
     *
     * Note of approximation: earth is not perfectly round,
     * at the 45th parallel (used in the test cases) corrective factors might be necessary.
     *
     * @return float
     */
    public function calculate(): float
    {
        $latitudeDelta = $this->lat2 - $this->lat1;
        $longitudeDelta = $this->lng2 - $this->lng1;

        $partial1 = pow(sin(($latitudeDelta) / 2), 2);
        $partial2 = cos($this->lat1) * cos($this->lat2) * pow(sin(($longitudeDelta) / 2), 2);

        return 2 * self::EARTH_RADIUS * asin(sqrt($partial1 + $partial2));
    }

    public function isCovered(): bool
    {
        return $this->calculate() <= self::COVERAGE_LIMIT;
    }
}
```

```php
<?php

namespace Task_2;

use PHPUnit\Framework\TestCase;

class HaversineTest extends TestCase
{
    private array $locations = [
        ['id' => 1010, 'zip_code' => '37169', 'lat' => 45.35, 'lng' => 10.84],
        ['id' => 1011, 'zip_code' => '37221', 'lat' => 45.44, 'lng' => 10.99],
        ['id' => 1012, 'zip_code' => '37229', 'lat' => 45.44, 'lng' => 11.00],
        ['id' => 1013, 'zip_code' => '37233', 'lat' => 45.43, 'lng' => 11.02],
    ];

    private array $shoppers = [
        ['id' => 'X1', 'lat' => 45.46, 'lng' => 11.03, 'enabled' => true],
        ['id' => 'X2', 'lat' => 45.46, 'lng' => 10.12, 'enabled' => true],
        ['id' => 'X3', 'lat' => 45.34, 'lng' => 10.81, 'enabled' => true],
        ['id' => 'X4', 'lat' => 45.76, 'lng' => 10.57, 'enabled' => true],
        ['id' => 'X5', 'lat' => 45.34, 'lng' => 10.63, 'enabled' => true],
        ['id' => 'X6', 'lat' => 45.42, 'lng' => 10.81, 'enabled' => true],
        ['id' => 'X7', 'lat' => 45.34, 'lng' => 10.94, 'enabled' => true],
    ];

    public function testCalculate(): void
    {
        foreach ($this->locations as $location) {
            foreach ($this->shoppers as $shopper) {
                $haversine = new Haversine($location['lat'], $location['lng'], $shopper['lat'], $shopper['lng']);
                $this->assertIsFloat($haversine->calculate());
                $this->assertTrue($haversine->calculate() > 0);
            }
        }
    }

    public function testIsCovered(): void
    {
        $totalCovered = 0;
        $shopperPercentage = [];

        foreach ($this->shoppers as $shopper) {
            $shopperCount = 0;

            foreach ($this->locations as $location) {
                $haversine = new Haversine($location['lat'], $location['lng'], $shopper['lat'], $shopper['lng']);
                if ($haversine->isCovered()) {
                    $totalCovered++;
                    $shopperCount++;
                }
            }

            $shopperPercentage[] = [
                'shopper_id' => $shopper['id'],
                'coverage' => $shopperCount / count($this->locations) * 100,
            ];
        }

        usort($shopperPercentage, function ($shopperA, $shopperB) {
            if ($shopperA['coverage'] < $shopperB['coverage']) {
                return 1;
            } elseif ($shopperA['coverage'] > $shopperB['coverage']) {
                return -1;
            } else {
                return 0;
            }
        });

        print_r(json_encode($shopperPercentage, JSON_PRETTY_PRINT));
        ob_flush();

        $this->assertEquals(6, $totalCovered);
    }
}
```