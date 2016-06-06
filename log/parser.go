//line parser.rl:1

// Copyright 2015 bs authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log

//go:generate bash -c "ragel -Z -G2 -o parser.go parser.rl && gofmt -s -w parser.go && gofmt -s -w parser.go"

//line parser.go:13
const lineparser_start int = 1
const lineparser_first_final int = 26
const lineparser_error int = 0

const lineparser_en_main int = 1

//line parser.rl:12
func parseLogLine(data []byte) [][]byte {
	entries := make([][]byte, 7)
	starts := make([]int, 7)
	cs, p, pe := 0, 0, len(data)
	eof := pe
	push := func(v int) {
		starts[v] = p
		entries[v] = nil
	}
	pop := func(v int) {
		entries[v] = data[starts[v]:p]
	}

//line parser.go:36
	{
		cs = lineparser_start
	}

//line parser.go:41
	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 26:
			goto st_case_26
		case 27:
			goto st_case_27
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 14:
			goto st_case_14
		case 15:
			goto st_case_15
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		case 28:
			goto st_case_28
		case 29:
			goto st_case_29
		case 30:
			goto st_case_30
		case 31:
			goto st_case_31
		case 32:
			goto st_case_32
		case 33:
			goto st_case_33
		case 34:
			goto st_case_34
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 23:
			goto st_case_23
		case 24:
			goto st_case_24
		case 25:
			goto st_case_25
		}
		goto st_out
	st_case_1:
		if data[p] == 60 {
			goto st2
		}
		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr2
		}
		goto st0
	tr2:
//line parser.rl:27
		push(0)
		goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line parser.go:146
		if data[p] == 62 {
			goto tr4
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st3
		}
		goto st0
	tr4:
//line parser.rl:27
		pop(0)
		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line parser.go:163
		if data[p] == 32 {
			goto st4
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st4
		}
		goto tr5
	tr5:
//line parser.rl:27
		push(1)
		goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line parser.go:180
		if data[p] == 32 {
			goto tr8
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr8
		}
		goto st5
	tr8:
//line parser.rl:27
		pop(1)
		goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line parser.go:197
		switch data[p] {
		case 32:
			goto st6
		case 58:
			goto tr11
		case 91:
			goto st0
		case 93:
			goto st0
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 13:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr11
				}
			case data[p] >= 9:
				goto st6
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr11
				}
			case data[p] >= 65:
				goto tr11
			}
		default:
			goto tr12
		}
		goto tr9
	tr9:
//line parser.rl:29
		push(4)
		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line parser.go:240
		switch data[p] {
		case 32:
			goto st0
		case 58:
			goto tr14
		case 91:
			goto tr15
		case 93:
			goto st0
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st0
		}
		goto st7
	tr14:
//line parser.rl:29
		pop(4)
		goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line parser.go:264
		switch data[p] {
		case 32:
			goto st9
		case 58:
			goto tr14
		case 91:
			goto tr15
		case 93:
			goto st0
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st9
		}
		goto st7
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		if data[p] == 32 {
			goto tr18
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr18
		}
		goto tr17
	tr17:
//line parser.rl:30
		push(6)
		goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line parser.go:300
		goto st26
	tr18:
//line parser.rl:30
		push(6)
		goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
//line parser.go:311
		if data[p] == 32 {
			goto tr18
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr18
		}
		goto tr17
	tr15:
//line parser.rl:29
		pop(4)
		goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line parser.go:328
		if 48 <= data[p] && data[p] <= 57 {
			goto tr19
		}
		goto st0
	tr19:
//line parser.rl:30
		push(5)
		goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line parser.go:342
		if data[p] == 93 {
			goto tr21
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st11
		}
		goto st0
	tr21:
//line parser.rl:30
		pop(5)
		goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line parser.go:359
		if data[p] == 58 {
			goto st13
		}
		goto st0
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		if data[p] == 32 {
			goto st9
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st9
		}
		goto st0
	tr11:
//line parser.rl:29
		push(3)
//line parser.rl:29
		push(4)
		goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line parser.go:387
		switch data[p] {
		case 32:
			goto tr23
		case 58:
			goto tr25
		case 91:
			goto tr15
		case 93:
			goto st0
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 13:
				if 45 <= data[p] && data[p] <= 46 {
					goto st14
				}
			case data[p] >= 9:
				goto tr23
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st14
				}
			case data[p] >= 65:
				goto st14
			}
		default:
			goto st14
		}
		goto st7
	tr23:
//line parser.rl:29
		pop(3)
		goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line parser.go:430
		switch data[p] {
		case 32:
			goto st15
		case 91:
			goto st0
		case 93:
			goto st0
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st15
		}
		goto tr9
	tr25:
//line parser.rl:29
		pop(4)
		goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line parser.go:452
		switch data[p] {
		case 32:
			goto tr27
		case 58:
			goto tr25
		case 91:
			goto tr15
		case 93:
			goto st0
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 13:
				if 45 <= data[p] && data[p] <= 46 {
					goto st14
				}
			case data[p] >= 9:
				goto tr27
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st14
				}
			case data[p] >= 65:
				goto st14
			}
		default:
			goto st14
		}
		goto st7
	tr27:
//line parser.rl:29
		pop(3)
		goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line parser.go:495
		switch data[p] {
		case 32:
			goto tr29
		case 91:
			goto tr17
		case 93:
			goto tr17
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr29
		}
		goto tr28
	tr28:
//line parser.rl:29
		push(4)
//line parser.rl:30
		push(6)
		goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line parser.go:519
		switch data[p] {
		case 32:
			goto st26
		case 58:
			goto tr42
		case 91:
			goto tr43
		case 93:
			goto st26
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st26
		}
		goto st28
	tr42:
//line parser.rl:29
		pop(4)
		goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
//line parser.go:543
		switch data[p] {
		case 32:
			goto st27
		case 58:
			goto tr42
		case 91:
			goto tr43
		case 93:
			goto st26
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st27
		}
		goto st28
	tr43:
//line parser.rl:29
		pop(4)
		goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
//line parser.go:567
		if 48 <= data[p] && data[p] <= 57 {
			goto tr45
		}
		goto st26
	tr45:
//line parser.rl:30
		push(5)
		goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
//line parser.go:581
		if data[p] == 93 {
			goto tr47
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st31
		}
		goto st26
	tr47:
//line parser.rl:30
		pop(5)
		goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
//line parser.go:598
		if data[p] == 58 {
			goto st33
		}
		goto st26
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		if data[p] == 32 {
			goto st27
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st27
		}
		goto st26
	tr29:
//line parser.rl:30
		push(6)
		goto st34
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
//line parser.go:624
		switch data[p] {
		case 32:
			goto tr29
		case 91:
			goto tr17
		case 93:
			goto tr17
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr29
		}
		goto tr28
	tr12:
//line parser.rl:28
		push(2)
//line parser.rl:29
		push(3)
//line parser.rl:29
		push(4)
		goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line parser.go:650
		switch data[p] {
		case 32:
			goto tr30
		case 58:
			goto tr25
		case 91:
			goto tr15
		case 93:
			goto st0
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 13:
				if 45 <= data[p] && data[p] <= 46 {
					goto st14
				}
			case data[p] >= 9:
				goto tr30
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st14
				}
			case data[p] >= 65:
				goto st14
			}
		default:
			goto st18
		}
		goto st7
	tr30:
//line parser.rl:29
		pop(3)
		goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line parser.go:693
		switch data[p] {
		case 32:
			goto st15
		case 91:
			goto st0
		case 93:
			goto st0
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr32
			}
		case data[p] >= 9:
			goto st15
		}
		goto tr9
	tr32:
//line parser.rl:29
		push(4)
		goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line parser.go:720
		switch data[p] {
		case 32:
			goto st0
		case 58:
			goto tr34
		case 91:
			goto tr15
		case 93:
			goto st0
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st20
			}
		case data[p] >= 9:
			goto st0
		}
		goto st7
	tr34:
//line parser.rl:29
		pop(4)
		goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line parser.go:749
		switch data[p] {
		case 32:
			goto st9
		case 58:
			goto tr14
		case 91:
			goto tr15
		case 93:
			goto st0
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st22
			}
		case data[p] >= 9:
			goto st9
		}
		goto st7
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		switch data[p] {
		case 32:
			goto st0
		case 58:
			goto tr36
		case 91:
			goto tr15
		case 93:
			goto st0
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st22
			}
		case data[p] >= 9:
			goto st0
		}
		goto st7
	tr36:
//line parser.rl:29
		pop(4)
		goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line parser.go:802
		switch data[p] {
		case 32:
			goto st9
		case 58:
			goto tr14
		case 91:
			goto tr15
		case 93:
			goto st0
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st24
			}
		case data[p] >= 9:
			goto st9
		}
		goto st7
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		switch data[p] {
		case 32:
			goto tr38
		case 58:
			goto tr14
		case 91:
			goto tr15
		case 93:
			goto st0
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto st24
			}
		case data[p] >= 9:
			goto tr38
		}
		goto st7
	tr38:
//line parser.rl:28
		pop(2)
		goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line parser.go:855
		switch data[p] {
		case 32:
			goto st25
		case 91:
			goto st0
		case 93:
			goto st0
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 13:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr11
				}
			case data[p] >= 9:
				goto st25
			}
		case data[p] > 58:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr11
				}
			case data[p] >= 65:
				goto tr11
			}
		default:
			goto tr11
		}
		goto tr9
	st_out:
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof26:
		cs = 26
		goto _test_eof
	_test_eof27:
		cs = 27
		goto _test_eof
	_test_eof10:
		cs = 10
		goto _test_eof
	_test_eof11:
		cs = 11
		goto _test_eof
	_test_eof12:
		cs = 12
		goto _test_eof
	_test_eof13:
		cs = 13
		goto _test_eof
	_test_eof14:
		cs = 14
		goto _test_eof
	_test_eof15:
		cs = 15
		goto _test_eof
	_test_eof16:
		cs = 16
		goto _test_eof
	_test_eof17:
		cs = 17
		goto _test_eof
	_test_eof28:
		cs = 28
		goto _test_eof
	_test_eof29:
		cs = 29
		goto _test_eof
	_test_eof30:
		cs = 30
		goto _test_eof
	_test_eof31:
		cs = 31
		goto _test_eof
	_test_eof32:
		cs = 32
		goto _test_eof
	_test_eof33:
		cs = 33
		goto _test_eof
	_test_eof34:
		cs = 34
		goto _test_eof
	_test_eof18:
		cs = 18
		goto _test_eof
	_test_eof19:
		cs = 19
		goto _test_eof
	_test_eof20:
		cs = 20
		goto _test_eof
	_test_eof21:
		cs = 21
		goto _test_eof
	_test_eof22:
		cs = 22
		goto _test_eof
	_test_eof23:
		cs = 23
		goto _test_eof
	_test_eof24:
		cs = 24
		goto _test_eof
	_test_eof25:
		cs = 25
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			switch cs {
			case 26, 27, 28, 29, 30, 31, 32, 33, 34:
//line parser.rl:30
				pop(6)
//line parser.go:928
			}
		}

	_out:
		{
		}
	}

//line parser.rl:33
	return entries
}
