package main

type KeyModifier int
const (
	KEY_MODIFIER_CONTROL KeyModifier = 1 << 0
	KEY_MODIFIER_ALT KeyModifier = 1 << 1
	KEY_MODIFIER_LOGO KeyModifier = 1 << 2
	KEY_MODIFIER_SHIFT KeyModifier = 1 << 3
	KEY_MODIFIER_ANY KeyModifier = ^0
)
