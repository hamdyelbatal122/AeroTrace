package main

import (
	"testing"

	"github.com/google/skylark"
)

func TestConvertHandlesCommonTypes(t *testing.T) {
	d := &skylark.Dict{}
	if err := d.SetKey(skylark.String("name"), skylark.String("nginx")); err != nil {
		t.Fatalf("set name: %v", err)
	}
	if err := d.SetKey(skylark.String("replicas"), skylark.MakeInt(3)); err != nil {
		t.Fatalf("set replicas: %v", err)
	}
	if err := d.SetKey(skylark.String("enabled"), skylark.True); err != nil {
		t.Fatalf("set enabled: %v", err)
	}

	got := convert(d)
	obj, ok := got.(map[string]interface{})
	if !ok {
		t.Fatalf("expected map[string]interface{}, got %T", got)
	}

	if obj["name"] != "nginx" {
		t.Fatalf("expected name nginx, got %#v", obj["name"])
	}

	replicas, ok := obj["replicas"].(int64)
	if !ok || replicas != 3 {
		t.Fatalf("expected replicas 3 as int64, got %#v (%T)", obj["replicas"], obj["replicas"])
	}

	enabled, ok := obj["enabled"].(bool)
	if !ok || !enabled {
		t.Fatalf("expected enabled true, got %#v (%T)", obj["enabled"], obj["enabled"])
	}
}

func TestConvertArray(t *testing.T) {
	list := skylark.NewList([]skylark.Value{
		skylark.String("a"),
		skylark.MakeInt(2),
		skylark.False,
	})

	got := convert(list)
	arr, ok := got.([]interface{})
	if !ok {
		t.Fatalf("expected []interface{}, got %T", got)
	}
	if len(arr) != 3 {
		t.Fatalf("expected len 3, got %d", len(arr))
	}
	if arr[0] != "a" {
		t.Fatalf("expected first item a, got %#v", arr[0])
	}
	if arr[1] != int64(2) {
		t.Fatalf("expected second item 2, got %#v", arr[1])
	}
	if arr[2] != false {
		t.Fatalf("expected third item false, got %#v", arr[2])
	}
}
