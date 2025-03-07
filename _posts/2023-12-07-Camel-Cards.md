---
title: Camel Cards
description: Advent of Code 2023 [Day 7]
layout: default
lang: en
tag: aoc23
prefetch:
  - adventofcode.com
---

Your all-expenses-paid trip turns out to be a one-way, five-minute ride in an airship. (At least it's a cool airship!) It drops you off at the edge of a vast desert and descends back to Island Island.

"Did you bring the parts?"

You turn around to see an Elf completely covered in white clothing, wearing goggles, and riding a large camel.

"Did you bring the parts?" she asks again, louder this time. You aren't sure what parts she's looking for; you're here to figure out why the sand stopped.

"The parts! For the sand, yes! Come with me; I will show you." She beckons you onto the camel.

After riding a bit across the sands of Desert Island, you can see what look like very large rocks covering half of the horizon. The Elf explains that the rocks are all along the part of Desert Island that is directly above Island Island, making it hard to even get there. Normally, they use big machines to move the rocks and filter the sand, but the machines have broken down because Desert Island recently stopped receiving the parts they need to fix the machines.

You've already assumed it'll be your job to figure out why the parts stopped when she asks if you can help. You agree automatically.

Because the journey will take a few days, she offers to teach you the game of Camel Cards. Camel Cards is sort of similar to poker except it's designed to be easier to play while riding a camel.

In Camel Cards, you get a list of hands, and your goal is to order them based on the strength of each hand. A hand consists of five cards labeled one of A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2. The relative strength of each card follows this order, where A is the highest and 2 is the lowest.

Every hand is exactly one type. From strongest to weakest, they are:

- Five of a kind, where all five cards have the same label: AAAAA
- Four of a kind, where four cards have the same label and one card has a different label: AA8AA
- Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
- Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
- Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
- One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
- High card, where all cards' labels are distinct: 23456

Hands are primarily ordered based on type; for example, every full house is stronger than any three of a kind.

If two hands have the same type, a second ordering rule takes effect. Start by comparing the first card in each hand. If these cards are different, the hand with the stronger first card is considered stronger. If the first card in each hand have the same label, however, then move on to considering the second card in each hand. If they differ, the hand with the higher second card wins; otherwise, continue with the third card in each hand, then the fourth, then the fifth.

So, 33332 and 2AAAA are both four of a kind hands, but 33332 is stronger because its first card is stronger. Similarly, 77888 and 77788 are both a full house, but 77888 is stronger because its third card is stronger (and both hands have the same first and second card).

To play Camel Cards, you are given a list of hands and their corresponding bid (your puzzle input). For example:

```
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
```

This example shows five hands; each hand is followed by its bid amount. Each hand wins an amount equal to its bid multiplied by its rank, where the weakest hand gets rank 1, the second-weakest hand gets rank 2, and so on up to the strongest hand. Because there are five hands in this example, the strongest hand will have rank 5 and its bid will be multiplied by 5.

So, the first step is to put the hands in order of strength:

- 32T3K is the only one pair and the other hands are all a stronger type, so it gets rank 1.
- KK677 and KTJJT are both two pair. Their first cards both have the same label, but the second card of KK677 is stronger (K vs T), so KTJJT gets rank 2 and KK677 gets rank 3.
- T55J5 and QQQJA are both three of a kind. QQQJA has a stronger first card, so it gets rank 5 and T55J5 gets rank 4.

Now, you can determine the total winnings of this set of hands by adding up the result of multiplying each hand's bid with its rank (765 * 1 + 220 * 2 + 28 * 3 + 684 * 4 + 483 * 5). So the total winnings in this example are 6440.

Find the rank of every hand in your set. What are the total winnings?

To make things a little more interesting, the Elf introduces one additional rule. Now, J cards are jokers - wildcards that can act like whatever card would make the hand the strongest type possible.

To balance this, J cards are now the weakest individual cards, weaker even than 2. The other cards stay in the same order: A, K, Q, T, 9, 8, 7, 6, 5, 4, 3, 2, J.

J cards can pretend to be whatever card is best for the purpose of determining hand type; for example, QJJQ2 is now considered four of a kind. However, for the purpose of breaking ties between two hands of the same type, J is always treated as J, not the card it's pretending to be: JKKK2 is weaker than QQQQ2 because J is weaker than Q.

Now, the above example goes very differently:

```
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
```

- 32T3K is still the only one pair; it doesn't contain any jokers, so its strength doesn't increase.
- KK677 is now the only two pair, making it the second-weakest hand.
- T55J5, KTJJT, and QQQJA are now all four of a kind! T55J5 gets rank 3, QQQJA gets rank 4, and KTJJT gets rank 5.

With the new joker rule, the total winnings in this example are 5905.

Using the new joker rule, find the rank of every hand in your set. What are the new total winnings?

```go
func SameScore(char rune, joker bool) uint {
	switch char {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		if joker {
			return 1
		}
		return 11
	case 'T':
		return 10
	case '9':
		return 9
	case '8':
		return 8
	case '7':
		return 7
	case '6':
		return 6
	case '5':
		return 5
	case '4':
		return 4
	case '3':
		return 3
	case '2':
		return 2
	default:
		return 0
	}
}

type Hand struct {
	Cards [5]rune
	Bid   uint
}

func (h *Hand) GetMatches(joker bool) uint {
	counts := make(map[rune]int)
	for _, r := range h.Cards {
		counts[r]++
	}

	if joker {
		for counts['J'] > 0 {
			vMaxNotJoker := -1
			var card rune
			for k, v := range counts {
				if k == 'J' {
					continue
				}
				if v > vMaxNotJoker {
					vMaxNotJoker = v
					card = k
				}
			}
			counts['J']--
			counts[card]++
		}
	}

	values := sort.IntSlice{}
	for _, v := range counts {
		values = append(values, v)
	}
	sort.Sort(sort.Reverse(values))

	switch values[0] {
	case 5:
		return 7
	case 4:
		return 6
	case 3:
		if values[1] == 2 {
			return 5
		}
		return 4
	case 2:
		if values[1] == 2 {
			return 3
		}
		return 2
	case 1:
		return 1
	default:
		return 0
	}
}

func (h *Hand) String() string {
	return fmt.Sprintf("Cards: [%s %s %s %s %s], Bid: %d",
		string(h.Cards[0]), string(h.Cards[1]),
		string(h.Cards[2]), string(h.Cards[3]),
		string(h.Cards[4]), h.Bid,
	)
}

type HandList struct {
	Hands []Hand
}

func (l *HandList) GetWinnings(joker bool) uint {
	sort.Slice(l.Hands, func(i, j int) bool {

		matches1 := l.Hands[i].GetMatches(joker)
		matches2 := l.Hands[j].GetMatches(joker)
		if matches1 < matches2 {
			return true
		}
		if matches1 > matches2 {
			return false
		}

		for c := 0; c < 5; c++ {
			card1 := SameScore(l.Hands[i].Cards[c], joker)
			card2 := SameScore(l.Hands[j].Cards[c], joker)
			if card1 < card2 {
				return true
			}
			if card1 > card2 {
				return false
			}
		}

		return true
	})

	winnings := uint(0)
	for i := 0; i < len(l.Hands); i++ {
		winnings += l.Hands[i].Bid * (uint(i) + 1)
	}
	return winnings
}

func (l *HandList) String() string {
	stringList := ""
	for pos, hand := range l.Hands {
		stringList += hand.String() + fmt.Sprintf(", Winnings: %d\n", hand.Bid*(uint(pos)+1))
	}
	return stringList
}

func TestHandList_GetWinnings(t *testing.T) {
	hl := &camelcards.HandList{
		Hands: []camelcards.Hand{
			{
				Cards: [5]rune{'3', '2', 'T', '3', 'K'},
				Bid:   765,
			},
			{
				Cards: [5]rune{'T', '5', '5', 'J', '5'},
				Bid:   684,
			},
			{
				Cards: [5]rune{'K', 'K', '6', '7', '7'},
				Bid:   28,
			},
			{
				Cards: [5]rune{'K', 'T', 'J', 'J', 'T'},
				Bid:   220,
			},
			{
				Cards: [5]rune{'Q', 'Q', 'Q', 'J', 'A'},
				Bid:   483,
			},
		},
	}

	t.Run("Example Test 2", func(t *testing.T) {
		l := hl
		if got := l.GetWinnings(true); got != 5905 {
			t.Errorf("HandList.GetWinnings() = %v, want 5905", got)
		}
	})

	t.Run("Example Test 1", func(t *testing.T) {
		l := hl
		if got := l.GetWinnings(false); got != 6440 {
			t.Errorf("HandList.GetWinnings() = %v, want 6440", got)
		}
	})
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	handList := camelcards.HandList{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		cards := []rune(line[0])

		bid, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatalln(err.Error())
		}

		handList.Hands = append(handList.Hands, camelcards.Hand{
			Cards: [5]rune(cards),
			Bid:   uint(bid),
		})
	}

	if err := scanner.Err(); err != nil {
		log.Println(err.Error())
	}

	points1 := handList.GetWinnings(false)
	dump1 := handList.String()
	points2 := handList.GetWinnings(true)
	dump2 := handList.String()
	log.Printf("\n---\n%s---\n%s---\nsum: %d, %d\n", dump1, dump2, points1, points2)
}
```

## Links

[If you're new to Advent of Code, it's an annual event that takes place throughout December, featuring a series of programming puzzles that get progressively more challenging as Christmas approaches.](https://adventofcode.com/2023/day/7)

<details>
	<summary>Click to show the input</summary>
	<pre>
8A7J7 301
QAAT7 677
J3K4K 622
KJJ62 577
49AAA 298
45585 855
33KKK 115
4Q777 438
7KK77 836
5T55T 397
85855 56
Q6Q38 157
AA8AA 85
32J33 293
KKQQA 247
888J4 944
2AJ2K 482
33777 338
KT434 696
K3K63 648
86866 136
93983 584
57857 489
5JJ2Q 76
82335 133
J25T4 559
9TQ2A 211
596J6 926
ATAAA 513
6KKKQ 277
AAA22 554
J2265 332
8Q3QQ 486
6735A 28
Q5555 595
J9888 262
5QQQ2 626
J7777 953
643TA 572
8579Q 99
23294 683
55J54 501
9JK93 567
64388 941
3J8T5 40
29K29 422
JQ4K2 401
AA6A3 78
2KK44 821
9AA2A 884
43434 386
J7A67 177
4JKKT 956
AA999 296
2A2T9 519
9T9KT 342
Q5J5A 19
QJK7A 925
AA9AA 337
4T2QT 751
77888 324
343QT 914
33229 576
J5Q5Q 169
22952 620
J4444 17
T9JTT 172
48888 729
28522 138
66363 302
68TTJ 778
5Q885 66
24KJ3 229
5A68K 731
A79A7 906
QQQAQ 698
J34AJ 109
TTTTJ 503
3J528 183
5A5AA 598
5AAAA 782
TT77J 968
24Q62 178
T6K7A 811
99788 53
494K6 560
7JQ87 327
Q9QQQ 597
496Q2 392
2252K 995
Q85TA 2
KT66Q 165
383T6 509
9997Q 724
4TT4T 368
35TQA 707
33534 24
KKKJ4 469
A9TAA 830
A445A 481
3Q63Q 192
AATAT 680
2225J 436
Q4869 471
A4JA4 330
5K35T 545
TK3A4 877
8K62J 596
9Q7T8 582
898KK 813
5AJJA 504
T5TKQ 289
6T88Q 359
A8AAJ 526
393J3 817
Q2Q22 212
Q8Q22 89
265T8 757
Q2584 807
T33TT 853
T22TA 391
46999 61
9AA9J 57
6T538 674
3T253 271
63TAJ 395
34943 323
6QQQQ 497
TA75K 938
4T467 141
36AQK 197
884A4 228
77277 453
AQ854 50
Q56QJ 456
T4TKT 319
K63Q4 30
79793 110
AJKQ6 421
22278 996
222TT 357
42J74 647
Q4QAT 634
66574 951
2KK22 446
Q65Q3 533
77887 267
58533 719
287K9 51
Q5656 199
564AQ 106
QJ579 77
9732T 46
JAK23 808
688K8 364
A993A 651
Q9QQ9 93
74777 929
QT722 723
76QAQ 621
8K836 573
77557 783
A5555 65
TTT6T 643
83595 933
J3J97 734
JTTA3 87
TTAQA 282
A555A 765
4TQTT 538
T777Q 763
T82J9 63
83359 213
AQQAQ 363
47AT9 311
96A9T 310
K5TK5 450
ATKK8 429
7K777 541
5TJAK 987
KAQ73 687
TTATA 874
J4TJ2 857
666KK 204
AA777 990
J98A8 946
7A95A 288
3K2K3 820
79K42 112
A65KT 300
596T8 266
J5664 585
75585 964
73AT2 491
2J2KT 375
6888T 480
A5765 992
KK555 224
AT2TA 514
TTTTQ 753
Q4348 273
KK66K 887
44JQ9 67
Q339Q 341
833Q8 574
K456T 766
J98KK 209
K4644 328
43935 79
8848K 380
K22A2 834
8654Q 703
4K2AT 96
K276J 863
8K64J 320
AAA7A 571
72747 950
4A4Q9 120
66636 814
33637 881
KK2AK 895
J8677 851
A48A8 430
56656 866
55355 976
66J77 370
3353J 896
56556 237
3KJ5Q 270
J22J2 261
AJ2AA 163
JJJJJ 517
48A28 819
K4443 799
TJTJ5 345
T3K9J 868
49AJ9 605
45455 741
KJJKK 754
63333 789
76677 989
AJ774 264
JKJ2A 531
TQQQT 809
Q6KJ6 411
KQK7Q 704
2A4Q5 410
68KK8 652
8K938 194
72J22 242
9J672 43
7Q6Q5 903
64666 95
J2372 798
A263T 781
JK598 909
55J92 625
A2567 343
4A7A7 174
8KA3J 68
JQTQ3 528
74444 127
QA5AT 306
T9K7T 254
75TT7 166
83388 325
888Q8 982
8Q7A6 84
2K7JA 663
T9959 937
3KA87 921
4KJ92 786
28A83 442
99T8J 699
TKKKK 457
25552 431
TTATK 862
QJQKK 664
QJQQ9 552
T4J77 206
Q7777 352
QJ747 534
9K9KK 656
44646 479
5T64Q 455
QQQQK 475
73478 827
Q3J33 975
Q655Q 930
7T7T7 231
53A72 409
3333A 833
29928 787
97J58 307
T2555 673
Q44A5 382
J5A57 144
7J447 131
8KK99 600
99555 349
43682 205
63259 518
AA3A8 591
J444J 58
52J23 772
5KK3K 660
5KT2K 329
54K6K 252
49J24 864
673Q5 780
TT848 333
227A5 424
JT9J5 675
7K9QJ 825
732Q7 461
24343 859
TJ333 604
T33AJ 149
KJ5KK 129
54555 249
25222 3
TQJA3 38
7Q7Q7 460
999J4 557
J6766 831
K4KKK 185
363K3 566
K888J 885
55855 488
97979 565
2597T 493
77456 233
Q9Q5Q 202
99K23 98
63763 7
TQ498 849
74626 590
K44KK 73
9AA97 498
Q6696 706
95555 697
69969 670
JQ294 203
24246 425
6QJ6A 16
AQAAA 916
6K6JK 313
9AK99 41
4T6J5 219
6Q6T6 549
96999 983
79977 522
57Q8A 732
AQ6AQ 156
284Q6 351
K88KJ 123
57A8A 606
4J97Q 592
5JTQ3 369
A467T 742
2878J 669
KKAJ6 668
T5TTT 570
8AA88 665
3238A 428
67286 979
88585 997
T6866 132
88T83 912
68888 832
A95TK 371
T8Q33 167
333K8 286
KK888 722
7TAK7 281
A7J77 216
9T85Q 969
54444 918
77822 151
9KJ2T 759
Q8483 536
9JJA9 788
J8Q8T 272
27KJQ 958
76666 671
5T932 182
887T4 628
6499T 607
447Q4 718
9736J 716
AQJAQ 611
T4933 690
3TJAJ 923
K7722 412
77737 22
775AQ 812
5QKQ5 362
448T8 114
A7A7A 502
4Q466 583
74744 749
7363J 939
88388 59
Q3KQT 861
42222 794
96669 294
54959 378
TTTTA 544
796JA 867
92KK9 238
9TJT9 848
8A2A2 309
687K3 515
KTKKT 610
J8659 210
22K4K 423
586Q9 55
AAA7K 599
33J73 198
K74KJ 936
55JJ5 943
Q6QJQ 801
5435J 137
QJ5QQ 553
AAA78 485
67676 404
34K33 376
8Q32J 326
A66KA 649
52K8Q 179
A32A2 561
JTQJQ 645
59KK7 11
6763Q 189
KKK88 426
8KTJ8 681
2222T 767
KK5J7 372
QQQ57 942
7A72J 904
2222J 467
7KQQQ 186
J3993 629
73437 952
8A37T 771
966J9 752
4KT29 710
82378 840
4QT72 785
3629T 934
JTK77 977
KKQ33 915
4242A 744
TQA9A 215
A6666 854
23333 170
TTT9T 587
J3377 34
6JJ66 835
A3AAA 897
96936 427
8JJJ4 686
32332 173
QJ868 640
67388 111
88988 613
4TJTT 435
T988T 702
K59K5 894
Q8Q88 688
5Q847 347
9J599 154
958K6 462
AAQQA 959
JKJKQ 846
7992T 72
9899A 510
22Q22 389
939J7 971
J4K4A 367
TT7TQ 816
KKKT3 297
97KK6 793
7QQQ7 824
ATQQA 102
J597K 797
78J88 473
87A87 42
828A7 527
K8777 107
8963J 949
29399 998
3TKA8 190
3TQ66 889
JJ5T5 985
J566T 568
T38K2 815
2J5KA 646
A8J32 184
622Q6 777
TK7QT 26
47766 790
2J5Q6 932
9K969 878
82235 64
3A3JK 152
49494 540
7JJJ7 350
39233 437
5466Q 239
Q3QQQ 961
A8J5Q 623
Q3333 739
4T396 90
25T2J 284
TK7T3 9
5555K 993
J2666 383
22292 250
J5T95 478
22T29 226
24242 126
99939 121
73373 414
78Q92 738
A5362 826
3KK2J 756
77877 291
T9TT9 879
3JJTT 145
3T363 524
6Q6QQ 619
Q2QQQ 8
8J58J 967
65KJK 398
KTA2T 464
T7JJJ 922
44888 948
2KJ2K 755
63858 353
49K4T 483
5833J 180
62822 105
99JK9 635
QAJAA 962
977T5 691
655KK 274
9966A 393
55J5T 303
284K5 589
27636 872
K29KK 774
K9JJ4 792
6JKTT 108
TK8TJ 745
T3A7K 400
9J3TQ 901
ATJKK 917
666A5 448
A5AT6 588
6T4J6 822
Q9A9A 260
594J3 241
Q3A34 279
A9999 220
7AQQ9 49
737Q3 999
JQQ8Q 882
78977 399
53333 689
2K2JQ 25
34848 800
2585T 334
5T799 908
69864 644
659TK 278
J4363 633
T9J97 101
QJ94T 615
363J6 373
J777J 379
99J93 966
TT26T 312
TA8T8 870
3J535 245
86A63 283
2KTK2 508
Q9Q2Q 162
59QJT 667
6JJKQ 71
55J25 200
89993 980
45474 581
Q4394 159
45697 346
KK99Q 496
J3J53 466
3QQ3Q 348
7843J 612
T733K 164
42589 818
Q9K36 232
QQ22Q 762
5AA3J 287
48KKK 449
7QQQJ 248
K5KJA 616
KQ387 965
Q4898 12
A3T9Q 227
K7KTT 269
JTJ22 768
85K88 875
6T666 624
97K7Q 586
7854J 804
T4TTT 609
8T44T 415
3T3KK 601
5AAAJ 258
8555K 828
4Q8JQ 593
JT823 396
J9K96 838
55565 632
7J77T 713
36353 578
5Q962 529
5KQ9T 181
82697 627
66654 88
7747J 919
A3AAJ 441
569KJ 805
TTQJ5 445
444T4 661
29959 37
62J26 955
Q4JT5 214
KQA4T 44
24424 758
2Q4AK 222
729A4 195
4Q35T 928
329T8 551
6T5T5 33
T25Q9 978
9Q9J9 407
A4KT7 911
27397 603
4244Q 490
JA255 384
2AAJ2 876
6222K 125
AA46A 360
J4343 672
JAJQ9 880
888J2 468
Q5K9J 268
A85Q9 387
J4A35 705
66A64 97
QAQQ8 700
55TJ4 31
88J58 650
693AK 679
259K2 806
9998J 176
6822T 235
9J699 913
7QQQQ 602
T7T4T 317
6226T 725
JJ9J9 316
55575 761
A6A66 796
K9T29 764
33334 608
KKK7K 893
99988 579
J8JJJ 259
4AJ89 390
9QQ84 14
9KJ3K 580
77736 433
KK22K 555
755J2 940
4JA9A 492
9Q555 484
6J687 542
99K53 910
6236J 36
9J999 54
6663J 454
85JJK 122
966J3 171
QJ2KK 747
KJA2A 676
788J5 858
78833 218
6KJ6J 308
T75AT 318
7JQ36 728
AJ8TA 354
96KTQ 899
52993 402
9K5AJ 898
Q7JQJ 748
5Q5QQ 48
A5858 148
85888 988
2TJ82 556
9AJK4 843
TJ26Q 406
2296K 356
22K22 18
KJ6KQ 253
T444T 117
6T257 715
AA444 113
39737 693
877QQ 521
TK454 735
38QA8 562
KT4A4 984
JKKKK 920
666J6 452
8JAJT 280
A62QJ 945
868TT 243
KQ66K 511
J3J33 512
T9K9K 160
66TT3 537
J3T48 234
44999 158
9876A 666
JJ73Q 546
85A3Q 20
3333K 750
6AJ73 381
Q3Q58 637
69644 654
J6669 641
8K2KK 733
33535 447
32K56 153
TA94Q 994
7979A 543
Q663Q 432
KQJ7Q 516
66K65 6
484T4 991
88387 15
96229 21
7AK62 499
88KK7 618
3K344 69
6A6KK 569
22J24 784
QKT53 837
7K2KK 1
87T95 810
2228Q 82
Q66AQ 548
J6Q66 366
5J42Q 465
8QQJ8 246
JAAKJ 630
5TT99 408
5T7Q6 779
768A7 631
22262 201
664AA 104
A96A6 128
T666K 5
3984K 684
3335K 208
79A77 891
K7K76 285
AQ8A6 678
36535 130
JT443 760
477T9 547
48A33 532
AA66A 321
56666 776
3AQKJ 888
K6Q68 83
2J727 902
3TT33 339
85Q6T 714
QT333 717
7T74T 335
J3A85 974
8J888 617
K974K 802
8777Q 947
A9Q6K 175
77757 75
7JA22 458
3K499 791
98TTJ 256
5555J 505
QTQQJ 70
T232T 506
Q9999 907
J4A88 74
A4K58 191
K3888 263
2QJ22 23
3368Q 155
QKKKQ 743
JK222 139
KQ84Q 905
T25T2 91
Q9TJJ 844
QQ8Q8 720
335KK 355
29T76 963
39AQ6 292
27222 276
88JJ8 388
52A22 134
7A95Q 86
7TTTA 314
9642K 187
A33A3 416
6Q66Q 92
93A23 403
77722 523
33J34 196
J8333 639
K8TJ5 736
222AA 116
JT3TT 970
87A65 225
82555 459
22JT2 255
Q2223 972
K63JA 143
8AQ86 653
JJ3J9 614
3A77A 365
KK9KT 47
5T959 385
7JKQ4 477
497T3 711
34A33 770
57557 869
777KJ 193
KKK5Q 892
9Q299 230
34J83 795
3KTTQ 692
66696 695
57472 960
77J26 841
TAT2T 685
33969 662
9Q393 295
QKQ88 142
TTT98 32
25555 374
265QK 924
JAAJA 440
9999K 257
346KT 500
JJ299 535
QJ4KA 135
7A422 305
A2T85 773
4T69J 52
73Q8A 463
KK6KK 420
8Q7QJ 251
QQK3K 168
66644 150
44424 444
84T53 299
AKAAK 240
6A952 35
Q5875 217
22888 829
38589 642
94848 417
32222 727
966Q9 27
JQ9QT 336
52K2T 886
J7787 721
AAQ3A 147
6K823 957
6668Q 839
AAA88 39
KJ33K 4
J5J5J 419
2A7T9 275
4A777 470
7KK7K 377
99992 418
93TA8 865
95T34 709
K34JA 322
94674 658
AJAA7 45
JQ22A 304
K83Q2 13
A6923 439
98977 331
689J6 769
99953 931
QQQJQ 94
J5AA5 871
TJTJT 413
ATA28 530
6T66T 564
KAKKK 873
T43QQ 657
TTTT2 563
9898J 118
QQQQ8 659
99997 236
88826 188
6KKK5 394
2QQJ2 495
8T864 900
888K8 655
28KQ4 361
33732 472
AA2A7 845
TTT7T 726
68A84 852
3KKKK 161
343Q2 358
75792 476
AAAJA 550
KK9A9 558
72569 737
4AT33 708
KKKQ9 119
8JJJ6 315
78A77 850
JJQJ7 140
3A5A5 740
78278 344
Q4Q44 103
44JA4 146
74T95 636
2A2AT 507
TK5Q5 474
7T982 712
AJ378 1000
75858 494
AAA99 746
72T94 981
TT77T 803
393K3 124
T8T95 265
AJA66 244
AA4TA 207
44777 638
68QT2 973
9K628 340
7595K 443
9A949 682
KJKTT 954
23T26 100
5KJ5T 434
66626 81
K33K3 856
J7555 525
JQJJ5 60
868J6 935
22299 823
J4445 539
9KKK7 890
JAA48 80
755Q5 520
4JQQJ 223
997T9 405
8A888 575
K369T 730
A5565 847
39JQ4 29
QQQJJ 694
86789 860
333J3 10
9A94A 221
J99JT 62
87888 701
AA6A2 290
T5QK4 775
35354 927
323Q3 451
25TA5 883
2954J 487
376J6 594
2K975 842
AAAA6 986
	</pre>
</details>
