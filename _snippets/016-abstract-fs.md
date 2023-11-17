---
title: Abstract Fs
---

```php
<?php

declare(strict_types=1);

namespace Task_1\Exceptions;

use Exception;
use Throwable;

class IllegalNameException extends Exception
{
    private const MESSAGE = 'Directory names consist only of English alphabet letters';

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

namespace Task_1\Exceptions;

use Exception;
use Throwable;

class InvalidRootException extends Exception
{
    private const MESSAGE = 'Root must be /';

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

namespace Task_1;

interface FilesystemAdapter
{
    public function cd(string $path): void;

    public function getCurrentPath(): string;
}
```

```php
<?php

declare(strict_types=1);

namespace Task_1;

use Task_1\Exceptions\IllegalNameException;
use Task_1\Exceptions\InvalidRootException;

class Path implements FilesystemAdapter
{
    private string $currentPath;
    private string $nextPath;

    public function __construct(string $path = '')
    {
        $this->cd($path);
    }

    /**
     * Change the current path of the class.
     *
     * Normalize the path in order to remove the "traverse to parent directory" characters.
     * Validates the given path for correctness (with basic regexes).
     *
     * @param string $path
     * @return void
     * @throws IllegalNameException
     * @throws InvalidRootException
     */
    public function cd(string $path): void
    {
        $this->nextPath = $this->getCurrentPath() . '/' . $this->normalizeSlashes($path);

        $this->validateNextPath();
        $this->normalizeNextPath();
        $this->validateNormalizedPath();

        $this->currentPath = $this->nextPath;
    }

    public function getCurrentPath(): string
    {
        return $this->currentPath ?? '';
    }

    private function validateNextPath(): void
    {
        if ($this->nextPath === '/') {
            return;
        }

        if (!str_starts_with($this->nextPath, '/')) {
            throw new InvalidRootException($this->nextPath);
        }

        if (preg_match('/^(\/([A-Za-z]+|\.\.))+\/?$/', $this->nextPath) !== 1) {
            throw new IllegalNameException($this->nextPath);
        }
    }

    private function validateNormalizedPath(): void
    {
        if ($this->nextPath === '/') {
            return;
        }

        if (!str_starts_with($this->nextPath, '/')) {
            throw new InvalidRootException($this->nextPath);
        }

        if (preg_match('/^(\/[A-Za-z]+)+$/', $this->nextPath) !== 1) {
            throw new IllegalNameException($this->nextPath);
        }
    }

    private function normalizeNextPath(): void
    {
        $count = 0;
        $this->nextPath = preg_replace('/\/[A-Za-z]+\/\.\./', '', $this->nextPath, 1, $count);

        if ($this->nextPath === '') {
            $this->nextPath = '/';
        }

        if ($count !== 0) {
            $this->normalizeNextPath();
        }
    }

    private function normalizeSlashes(string $path): string
    {
        return ltrim(rtrim($path, '/'), '/');
    }
}
```

```php
<?php

namespace Task_1;

use PHPUnit\Framework\TestCase;

class PathTest extends TestCase
{

    public function testGiven(): void
    {
        $path = new Path('/a/b/c/d');
        $path->cd('../x');
        $this->assertEquals('/a/b/c/x', $path->getCurrentPath());
    }

    public function testRoot(): void
    {
        $path = new Path();
        $this->assertEquals('/', $path->getCurrentPath());

        $path = new Path('');
        $this->assertEquals('/', $path->getCurrentPath());

        $path = new Path('/');
        $this->assertEquals('/', $path->getCurrentPath());
    }

    public function testEscape(): void
    {
        $this->expectException('Task_1\Exceptions\IllegalNameException');
        new Path('/../../');
    }

    public function testMultipleTraversals(): void
    {
        $path = new Path('/aa/bb/../cc/');
        $path->cd('../xx/dd/');
        $this->assertEquals('/aa/xx/dd', $path->getCurrentPath());
    }
}
```
