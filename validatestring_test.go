package uuidv7_test

import (
	"testing"

	"errors"

	"github.com/reiver/go-uuidv7"
)

func TestValidateString(t *testing.T) {
	tests := []struct{
		String   string
		Expected error
	}{
		{
			String: "019E0BC7-407B-7CD1-AA11-BF34CBBCC553",
			Expected: nil,

		},
		{
			String: "019E0bC7-407b-7Cd1-Aa11-Bf34CbBcC553",
			Expected: nil,
		},
		{
			String: "019e0bc7-407b-7cd1-aa11-bf34cbbcc553",
			Expected: nil,
		},



		{
			String: "019e0d44-0b23-7a3f-b3af-7a54d92cf850",
			Expected: nil,
		},



		{
			String: "019e1a5e-b6d4-757c-8890-9d2969f29d7e",
			Expected: nil,
		},



		{
			String: "fedcba98-7654-7321-8024-68acefdb9753",
			Expected: nil,
		},



		{
			String: "00000000-0000-7000-8000-000000000000",
			Expected: nil,
		},
		{
			String: "00000000-0000-7000-9000-000000000000",
			Expected: nil,
		},
		{
			String: "00000000-0000-7000-A000-000000000000",
			Expected: nil,
		},
		{
			String: "00000000-0000-7000-B000-000000000000",
			Expected: nil,
		},
		{
			String: "00000000-0000-7000-a000-000000000000",
			Expected: nil,
		},
		{
			String: "00000000-0000-7000-b000-000000000000",
			Expected: nil,
		},









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
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                     v
			String: "ed7ba470-8e54-465e-825c-99712043e01c",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                     v
			String: "eD7bA470-8e54-465E-825c-99712043E01c",
			Expected: uuidv7.ErrNotUUIDv7,
		},



		{
			//                     v
			String: "00000000-0000-0000-8000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                     v
			String: "00000000-0000-1000-8000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                     v
			String: "00000000-0000-2000-8000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                     v
			String: "00000000-0000-3000-8000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                     v
			String: "00000000-0000-4000-8000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                     v
			String: "00000000-0000-5000-8000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                     v
			String: "00000000-0000-6000-8000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},

		{
			//                     v
			String: "00000000-0000-8000-8000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                     v
			String: "00000000-0000-9000-8000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                     v
			String: "00000000-0000-A000-8000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                     v
			String: "00000000-0000-a000-8000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                     v
			String: "00000000-0000-B000-8000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                     v
			String: "00000000-0000-b000-8000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                     v
			String: "00000000-0000-C000-8000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                     v
			String: "00000000-0000-c000-8000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                     v
			String: "00000000-0000-D000-8000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                     v
			String: "00000000-0000-d000-8000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                     v
			String: "00000000-0000-E000-8000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                     v
			String: "00000000-0000-e000-8000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                     v
			String: "00000000-0000-F000-8000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                     v
			String: "00000000-0000-f000-8000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},

		{
			//                          v
			String: "00000000-0000-7000-0000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                          v
			String: "00000000-0000-7000-1000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                          v
			String: "00000000-0000-7000-2000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                          v
			String: "00000000-0000-7000-3000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                          v
			String: "00000000-0000-7000-4000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                          v
			String: "00000000-0000-7000-5000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                          v
			String: "00000000-0000-7000-6000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                          v
			String: "00000000-0000-7000-7000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},

		{
			//                          v
			String: "00000000-0000-7000-C000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                          v
			String: "00000000-0000-7000-c000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                          v
			String: "00000000-0000-7000-D000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                          v
			String: "00000000-0000-7000-d000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                          v
			String: "00000000-0000-7000-E000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                          v
			String: "00000000-0000-7000-e000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                          v
			String: "00000000-0000-7000-F000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
		{
			//                          v
			String: "00000000-0000-7000-f000-000000000000",
			Expected: uuidv7.ErrNotUUIDv7,
		},
	}

	for testNumber, test := range tests {
		err := uuidv7.ValidateString(test.String)
		switch test.Expected {
		case nil:
			if nil != err {
				t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
				t.Logf("ERROR: %s", err)
				t.Logf("STRING: %s", test.String)
				continue
			}
		default:
			if !errors.Is(err, test.Expected) {
				t.Errorf("For test #%d, the actual error is not what was expected.", testNumber)
				t.Logf("EXPECTED-ERROR: %s", test.Expected)
				t.Logf("ACTUAL-ERROR:   %s", err)
				t.Logf("STRING: %s", test.String)
				continue
			}
		}

	}
}
