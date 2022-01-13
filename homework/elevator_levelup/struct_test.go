package main

import (
	"reflect"
	"testing"
)

func Test_delItem(t *testing.T) {
	type args struct {
		slice []int
		item  int
	}
	test := args{
		[]int{1, 2, 3, 4, 5},
		3,
	}
	if got := delItem(test.slice, test.item); !reflect.DeepEqual(got, []int{1, 2, 4, 5}) {
		t.Errorf("delItem() = %v, want %v", got, []int{1, 2, 4, 5})
	}
}

func Test_elevator_add(t *testing.T) {
	e := elevator{
		floor:    1,
		target:   nil,
		maxFloor: 1,
		minFloor: 1,
	}
	e.add(1, 2, 3)
	if got := e; !reflect.DeepEqual(got, elevator{
		floor:    1,
		target:   []int{1, 2, 3},
		maxFloor: 3,
		minFloor: 1,
	}) {
		t.Errorf("e = %v, want %v", got, elevator{
			floor:    1,
			target:   []int{1, 2, 3},
			maxFloor: 3,
			minFloor: 1,
		})
	}
}

func Test_elevator_arrive(t *testing.T) {
	type fields struct {
		floor    int
		target   []int
		maxFloor int
		minFloor int
	}
	type args struct {
		f int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &elevator{
				floor:    tt.fields.floor,
				target:   tt.fields.target,
				maxFloor: tt.fields.maxFloor,
				minFloor: tt.fields.minFloor,
			}
			if err := e.arrive(tt.args.f); (err != nil) != tt.wantErr {
				t.Errorf("arrive() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_elevator_down(t *testing.T) {
	type fields struct {
		floor    int
		target   []int
		maxFloor int
		minFloor int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &elevator{
				floor:    tt.fields.floor,
				target:   tt.fields.target,
				maxFloor: tt.fields.maxFloor,
				minFloor: tt.fields.minFloor,
			}
			if err := e.down(); (err != nil) != tt.wantErr {
				t.Errorf("down() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_elevator_up(t *testing.T) {
	type fields struct {
		floor    int
		target   []int
		maxFloor int
		minFloor int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &elevator{
				floor:    tt.fields.floor,
				target:   tt.fields.target,
				maxFloor: tt.fields.maxFloor,
				minFloor: tt.fields.minFloor,
			}
			if err := e.up(); (err != nil) != tt.wantErr {
				t.Errorf("up() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
