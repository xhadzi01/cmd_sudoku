package main

import "errors"

const (
	BoardSideSize int = 9
)

type Slot interface {
	IsPreset() bool
	IsEmpty() bool
	IsSelected() bool
	SetSelected(bool)
	Value() int
	SetValue(int) error
}

type SelectableSlot struct {
	selected bool
}

func (slot *SelectableSlot) IsSelected() bool {
	return slot.selected
}
func (slot *SelectableSlot) SetSelected(isSelected bool) {
	slot.selected = isSelected
}

type FillableSlot struct {
	SelectableSlot
	value *int
}

func (slot *FillableSlot) IsPreset() bool {
	return false
}

func (slot *FillableSlot) IsEmpty() bool {
	return nil == slot.value
}

func (slot *FillableSlot) Value() int {
	if nil == slot.value {
		panic("Value not set")
	}
	return *slot.value
}

func (slot *FillableSlot) SetValue(newValue int) error {
	if newValue == 0 {
		slot.value = nil
		return nil
	}

	slot.value = &newValue
	return nil
}

func NewFillableSlot() Slot {
	return &FillableSlot{
		SelectableSlot: SelectableSlot{
			selected: false,
		},
		value: nil,
	}
}

type PresetSlot struct {
	SelectableSlot
	value int
}

func (slot *PresetSlot) IsPreset() bool {
	return true
}

func (slot *PresetSlot) IsEmpty() bool {
	return false
}

func (slot *PresetSlot) Value() int {
	return slot.value
}

func (slot *PresetSlot) SetValue(newValue int) error {
	return errors.New("could not set preset field")
}

func NewPresetSlot(value int) Slot {
	return &PresetSlot{
		SelectableSlot: SelectableSlot{
			selected: false,
		},
		value: value,
	}
}

type PlayingBoard [BoardSideSize * BoardSideSize]Slot
