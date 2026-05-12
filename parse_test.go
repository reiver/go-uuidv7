package uuidv7_test

import (
	"testing"

	"errors"

	"github.com/reiver/go-uuidv7"
)

func TestParse(t *testing.T) {
	tests := []struct{
		String   string
		Expected [16]byte
	}{
		{
			String: "019E0BC7-407B-7CD1-AA11-BF34CBBCC553",
			Expected: [16]byte{0x01,0x9e,0x0b,0xc7,0x40,0x7b,0x7c,0xd1,0xaa,0x11,0xbf,0x34,0xcb,0xbc,0xc5,0x53},
		},
		{
			String: "019E0bC7-407b-7Cd1-Aa11-Bf34CbBcC553",
			Expected: [16]byte{0x01,0x9e,0x0b,0xc7,0x40,0x7b,0x7c,0xd1,0xaa,0x11,0xbf,0x34,0xcb,0xbc,0xc5,0x53},
		},
		{
			String: "019e0bc7-407b-7cd1-aa11-bf34cbbcc553",
			Expected: [16]byte{0x01,0x9e,0x0b,0xc7,0x40,0x7b,0x7c,0xd1,0xaa,0x11,0xbf,0x34,0xcb,0xbc,0xc5,0x53},
		},



		{
			String: "019e0d44-0b23-7a3f-b3af-7a54d92cf850",
			Expected: [16]byte{0x01,0x9e,0x0d,0x44,0x0b,0x23,0x7a,0x3f,0xb3,0xaf,0x7a,0x54,0xd9,0x2c,0xf8,0x50},
		},



		{
			String: "019e1a5e-b6d4-757c-8890-9d2969f29d7e",
			Expected: [16]byte{0x01,0x9e,0x1a,0x5e,0xb6,0xd4,0x75,0x7c,0x88,0x90,0x9d,0x29,0x69,0xf2,0x9d,0x7e},
		},



		{
			String: "fedcba98-7654-7321-8024-68acefdb9753",
			Expected: [16]byte{0xfe,0xdc,0xba,0x98,0x76,0x54,0x73,0x21,0x80,0x24,0x68,0xac,0xef,0xdb,0x97,0x53},
		},



		{
			String: "00000000-0000-7000-8000-000000000000",
			Expected: [16]byte{0x00,0x00,0x00,0x00,0x00,0x00,0x70,0x00,0x80,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
		},
		{
			String: "00000000-0000-7000-9000-000000000000",
			Expected: [16]byte{0x00,0x00,0x00,0x00,0x00,0x00,0x70,0x00,0x90,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
		},
		{
			String: "00000000-0000-7000-A000-000000000000",
			Expected: [16]byte{0x00,0x00,0x00,0x00,0x00,0x00,0x70,0x00,0xa0,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
		},
		{
			String: "00000000-0000-7000-B000-000000000000",
			Expected: [16]byte{0x00,0x00,0x00,0x00,0x00,0x00,0x70,0x00,0xb0,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
		},
		{
			String: "00000000-0000-7000-a000-000000000000",
			Expected: [16]byte{0x00,0x00,0x00,0x00,0x00,0x00,0x70,0x00,0xa0,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
		},
		{
			String: "00000000-0000-7000-b000-000000000000",
			Expected: [16]byte{0x00,0x00,0x00,0x00,0x00,0x00,0x70,0x00,0xb0,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
		},
	}

	for i:=0; i<64; i++ {
		uuidBytes  := uuidv7.Generate()
		uuidString := uuidv7.String(uuidBytes)

		test := struct{
			String   string
			Expected [16]byte
		}{
			String:   uuidString,
			Expected: uuidBytes,
		}

		tests = append(tests, test)
	}

	for testNumber, test := range tests {
		actual, err := uuidv7.Parse(test.String)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("STRING: %s", test.String)
			continue
		}

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual parsed UUID is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			t.Logf("STRING: %s", test.String)
			continue
		}

		if !uuidv7.Is(actual) {
			t.Errorf("For test #%d, expected a verion 7 UUID but did not actually get one.", testNumber)
			t.Logf("ACTUAL:   %#v", actual)
			t.Logf("STRING: %s", test.String)
			continue
		}

		if !uuidv7.IsString(test.String) {
			t.Errorf("For test #%d, the string should have been a verion 7 UUID but did actually wasn't.", testNumber)
			t.Logf("STRING: %s", test.String)
			continue
		}
	}
}

func TestParse_fail(t *testing.T) {
	tests := []struct{
		String   string
		Expected error
	}{
		{
			String: "",
			Expected: uuidv7.ErrLengthWrong,
		},
		{
			String: "ED7BA4708E54465E825C99712043E01C",
			Expected: uuidv7.ErrLengthWrong,
		},
		{
			String: "ED7BA470-8E54-465E-825C-99712043E01Cf",
			Expected: uuidv7.ErrLengthWrong,
		},



		{
			String: "000000000000000000000000000000000000",
			Expected: uuidv7.ErrHyphenMissing,
		},
		{
			String: "111111111111111111111111111111111111",
			Expected: uuidv7.ErrHyphenMissing,
		},
		{
			String: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
			Expected: uuidv7.ErrHyphenMissing,
		},
		{
			String: "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF",
			Expected: uuidv7.ErrHyphenMissing,
		},
		{
			String: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			Expected: uuidv7.ErrHyphenMissing,
		},
		{
			String: "ffffffffffffffffffffffffffffffffffff",
			Expected: uuidv7.ErrHyphenMissing,
		},



		{
			String: "------------------------------------",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			String: "--------------7----8----------------",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			String: "--------------7----9----------------",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			String: "--------------7----A----------------",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			String: "--------------7----B----------------",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			String: "--------------7----a----------------",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			String: "--------------7----b----------------",
			Expected: uuidv7.ErrNotUUID,
		},



		{
			String: "4-----------------------------------",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			String: "4-------------7----8----------------",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			String: "4-------------7----9----------------",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			String: "4-------------7----A----------------",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			String: "4-------------7----B----------------",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			String: "4-------------7----a----------------",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			String: "4-------------7----b----------------",
			Expected: uuidv7.ErrNotUUID,
		},



		{
			//       v
			String: "-19e0d44-0b23-7a3f-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//        v
			String: "0-9e0d44-0b23-7a3f-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//         v
			String: "01-e0d44-0b23-7a3f-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//          v
			String: "019-0d44-0b23-7a3f-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//           v
			String: "019e-d44-0b23-7a3f-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//            v
			String: "019e0-44-0b23-7a3f-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//             v
			String: "019e0d-4-0b23-7a3f-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//              v
			String: "019e0d4--0b23-7a3f-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},

		{
			//                v
			String: "019e0d44--b23-7a3f-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                 v
			String: "019e0d44-0-23-7a3f-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                  v
			String: "019e0d44-0b-3-7a3f-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                   v
			String: "019e0d44-0b2--7a3f-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},

		{
			//                     v
			String: "019e0d44-0b23--a3f-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                      v
			String: "019e0d44-0b23-7-3f-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                       v
			String: "019e0d44-0b23-7a-f-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                        v
			String: "019e0d44-0b23-7a3--b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},

		{
			//                          v
			String: "019e0d44-0b23-7a3f--3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                           v
			String: "019e0d44-0b23-7a3f-b-af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                            v
			String: "019e0d44-0b23-7a3f-b3-f-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                             v
			String: "019e0d44-0b23-7a3f-b3a--7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},

		{
			//                               v
			String: "019e0d44-0b23-7a3f-b3af--a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                                v
			String: "019e0d44-0b23-7a3f-b3af-7-54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                                 v
			String: "019e0d44-0b23-7a3f-b3af-7a-4d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                                  v
			String: "019e0d44-0b23-7a3f-b3af-7a5-d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                                   v
			String: "019e0d44-0b23-7a3f-b3af-7a54-92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                                    v
			String: "019e0d44-0b23-7a3f-b3af-7a54d-2cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                                     v
			String: "019e0d44-0b23-7a3f-b3af-7a54d9-cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                                      v
			String: "019e0d44-0b23-7a3f-b3af-7a54d92-f850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                                       v
			String: "019e0d44-0b23-7a3f-b3af-7a54d92c-850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                                        v
			String: "019e0d44-0b23-7a3f-b3af-7a54d92cf-50",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                                         v
			String: "019e0d44-0b23-7a3f-b3af-7a54d92cf8-0",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                                          v
			String: "019e0d44-0b23-7a3f-b3af-7a54d92cf85-",
			Expected: uuidv7.ErrNotUUID,
		},



		{
			//       v
			String: "g19e0d44-0b23-7a3f-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//        v
			String: "0g9e0d44-0b23-7a3f-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//         v
			String: "01ge0d44-0b23-7a3f-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//          v
			String: "019g0d44-0b23-7a3f-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//           v
			String: "019egd44-0b23-7a3f-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//            v
			String: "019e0g44-0b23-7a3f-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//             v
			String: "019e0dg4-0b23-7a3f-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//              v
			String: "019e0d4g-0b23-7a3f-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},

		{
			//                v
			String: "019e0d44-gb23-7a3f-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                 v
			String: "019e0d44-0g23-7a3f-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                  v
			String: "019e0d44-0bg3-7a3f-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                   v
			String: "019e0d44-0b2g-7a3f-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},

		{
			//                     v
			String: "019e0d44-0b23-ga3f-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                      v
			String: "019e0d44-0b23-7g3f-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                       v
			String: "019e0d44-0b23-7agf-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                        v
			String: "019e0d44-0b23-7a3g-b3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},

		{
			//                          v
			String: "019e0d44-0b23-7a3f-g3af-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                           v
			String: "019e0d44-0b23-7a3f-bgaf-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                            v
			String: "019e0d44-0b23-7a3f-b3gf-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                             v
			String: "019e0d44-0b23-7a3f-b3ag-7a54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},

		{
			//                               v
			String: "019e0d44-0b23-7a3f-b3af-ga54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                                v
			String: "019e0d44-0b23-7a3f-b3af-7g54d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                                 v
			String: "019e0d44-0b23-7a3f-b3af-7ag4d92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                                  v
			String: "019e0d44-0b23-7a3f-b3af-7a5gd92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                                   v
			String: "019e0d44-0b23-7a3f-b3af-7a54g92cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                                    v
			String: "019e0d44-0b23-7a3f-b3af-7a54dg2cf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                                     v
			String: "019e0d44-0b23-7a3f-b3af-7a54d9gcf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                                      v
			String: "019e0d44-0b23-7a3f-b3af-7a54d92gf850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                                       v
			String: "019e0d44-0b23-7a3f-b3af-7a54d92cg850",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                                        v
			String: "019e0d44-0b23-7a3f-b3af-7a54d92cfg50",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                                         v
			String: "019e0d44-0b23-7a3f-b3af-7a54d92cf8g0",
			Expected: uuidv7.ErrNotUUID,
		},
		{
			//                                          v
			String: "019e0d44-0b23-7a3f-b3af-7a54d92cf85g",
			Expected: uuidv7.ErrNotUUID,
		},



		{
			//                     v
			String: "ED7BA470-8E54-465E-825C-99712043E01C",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                     v
			String: "ed7ba470-8e54-465e-825c-99712043e01c",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                     v
			String: "eD7bA470-8e54-465E-825c-99712043E01c",
			Expected: uuidv7.ErrNotVersion7UUID,
		},



		{
			//                     v
			String: "00000000-0000-0000-8000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                     v
			String: "00000000-0000-1000-8000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                     v
			String: "00000000-0000-2000-8000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                     v
			String: "00000000-0000-3000-8000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                     v
			String: "00000000-0000-4000-8000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                     v
			String: "00000000-0000-5000-8000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                     v
			String: "00000000-0000-6000-8000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},

		{
			//                     v
			String: "00000000-0000-8000-8000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                     v
			String: "00000000-0000-9000-8000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                     v
			String: "00000000-0000-A000-8000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                     v
			String: "00000000-0000-a000-8000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                     v
			String: "00000000-0000-B000-8000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                     v
			String: "00000000-0000-b000-8000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                     v
			String: "00000000-0000-C000-8000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                     v
			String: "00000000-0000-c000-8000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                     v
			String: "00000000-0000-D000-8000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                     v
			String: "00000000-0000-d000-8000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                     v
			String: "00000000-0000-E000-8000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                     v
			String: "00000000-0000-e000-8000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                     v
			String: "00000000-0000-F000-8000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                     v
			String: "00000000-0000-f000-8000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},

		{
			//                          v
			String: "00000000-0000-7000-0000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                          v
			String: "00000000-0000-7000-1000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                          v
			String: "00000000-0000-7000-2000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                          v
			String: "00000000-0000-7000-3000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                          v
			String: "00000000-0000-7000-4000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                          v
			String: "00000000-0000-7000-5000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                          v
			String: "00000000-0000-7000-6000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                          v
			String: "00000000-0000-7000-7000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},

		{
			//                          v
			String: "00000000-0000-7000-C000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                          v
			String: "00000000-0000-7000-c000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                          v
			String: "00000000-0000-7000-D000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                          v
			String: "00000000-0000-7000-d000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                          v
			String: "00000000-0000-7000-E000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                          v
			String: "00000000-0000-7000-e000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                          v
			String: "00000000-0000-7000-F000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},
		{
			//                          v
			String: "00000000-0000-7000-f000-000000000000",
			Expected: uuidv7.ErrNotVersion7UUID,
		},

	}

	for testNumber, test := range tests {
		_, actual := uuidv7.Parse(test.String)
		if nil == actual {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("STRING: %s", test.String)
			continue
		}

		expected := test.Expected

		if !errors.Is(actual, expected) {
			t.Errorf("For test #%d, the actually error is not what was expected.", testNumber)
			t.Logf("EXPECTED: %s", expected)
			t.Logf("ACTUAL:   %s", actual)
			t.Logf("STRING: %s", test.String)
			continue
		}

		if uuidv7.IsString(test.String) {
			t.Errorf("For test #%d, the string should not have been a version 7 UUID but it actually was.", testNumber)
			t.Logf("STRING: %s", test.String)
			continue
		}

	}
}
