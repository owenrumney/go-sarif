package test

import "testing"

func Test_new_result_on_run(t *testing.T) {
	given, _, _ := newRunTest(t)

	given.a_result_is_added()

}
