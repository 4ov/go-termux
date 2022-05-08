package termux

import (
	"encoding/json"
)

type DialogConfirmOptions struct {
	// text hint
	Hint string `arg:"-i"`
	// set title of dialog
	Title string `arg:"-t"`
}

type DialogConfirmResult struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func DialogConfirm(options DialogConfirmOptions) (DialogConfirmResult, error) {
	args := []string{"confirm"}
	args = append(args, ReadyArgs(options)...)
	output, err := CallCommand("termux-dialog", args...)
	if err != nil {
		return DialogConfirmResult{}, err
	}
	result := DialogConfirmResult{}
	err = json.Unmarshal(output, &result)
	if err != nil {
		return DialogConfirmResult{}, err
	}
	return result, nil

}

type DialogCheckboxOptions struct {
	// text hint
	Values []string `arg:"-v" split:","`
	// set title of dialog
	Title string `arg:"-t"`
}

type DialogCheckboxResult struct {
	Code   int    `json:"code"`
	Text   string `json:"text"`
	Values []struct {
		Index int    `json:"index"`
		Value string `json:"value"`
	}
}

func DialogCheckBox(options DialogCheckboxOptions) (DialogCheckboxResult, error) {
	args := []string{"checkbox"}
	args = append(args, ReadyArgs(options)...)
	output, err := CallCommand("termux-dialog", args...)
	if err != nil {
		return DialogCheckboxResult{}, err
	}
	result := DialogCheckboxResult{}
	err = json.Unmarshal(output, &result)
	if err != nil {
		return DialogCheckboxResult{}, err
	}
	return result, nil
}
