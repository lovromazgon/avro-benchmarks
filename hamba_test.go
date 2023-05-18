package avro_benchmarks

import (
	"testing"

	"github.com/hamba/avro/v2"
	"github.com/stretchr/testify/assert"
)

func TestHambaDecode(t *testing.T) {
	schema := avro.MustParse(Schema)

	superhero := Superhero{}
	err := avro.Unmarshal(schema, Payload, &superhero)

	want := Superhero{
		ID:            234765,
		AffiliationID: 9867,
		Name:          "Wolverine",
		Life:          85.25,
		Energy:        32.75,
		Powers: []*Superpower{
			{ID: 2345, Name: "Bone Claws", Damage: 5, Energy: 1.15, Passive: false},
			{ID: 2346, Name: "Regeneration", Damage: -2, Energy: 0.55, Passive: true},
			{ID: 2347, Name: "Adamant skeleton", Damage: -10, Energy: 0, Passive: true},
		},
	}
	assert.NoError(t, err)
	assert.Equal(t, want, superhero)
}

func BenchmarkHambaDecode(b *testing.B) {
	schema := avro.MustParse(Schema)

	superhero := Superhero{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = avro.Unmarshal(schema, Payload, &superhero)
	}
}

func TestHambaEncode(t *testing.T) {
	schema := avro.MustParse(Schema)

	superhero := Superhero{
		ID:            234765,
		AffiliationID: 9867,
		Name:          "Wolverine",
		Life:          85.25,
		Energy:        32.75,
		Powers: []*Superpower{
			{ID: 2345, Name: "Bone Claws", Damage: 5, Energy: 1.15, Passive: false},
			{ID: 2346, Name: "Regeneration", Damage: -2, Energy: 0.55, Passive: true},
			{ID: 2347, Name: "Adamant skeleton", Damage: -10, Energy: 0, Passive: true},
		},
	}

	b, err := avro.Marshal(schema, &superhero)

	assert.NoError(t, err)
	assert.Equal(t, Payload, b)
}

func BenchmarkHambaEncode(b *testing.B) {
	schema := avro.MustParse(Schema)

	superhero := Superhero{
		ID:            234765,
		AffiliationID: 9867,
		Name:          "Wolverine",
		Life:          85.25,
		Energy:        32.75,
		Powers: []*Superpower{
			{ID: 2345, Name: "Bone Claws", Damage: 5, Energy: 1.15, Passive: false},
			{ID: 2346, Name: "Regeneration", Damage: -2, Energy: 0.55, Passive: true},
			{ID: 2347, Name: "Adamant skeleton", Damage: -10, Energy: 0, Passive: true},
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = avro.Marshal(schema, &superhero)
	}
}

func TestHambaDecodeMap(t *testing.T) {
	schema := avro.MustParse(Schema)

	superhero := map[string]interface{}{}
	err := avro.Unmarshal(schema, Payload, &superhero)

	want := map[string]interface{}{
		"id":             234765,
		"affiliation_id": 9867,
		"name":           "Wolverine",
		"life":           float32(85.25),
		"energy":         float32(32.75),
		"powers": []interface{}{
			map[string]interface{}{"id": 2345, "name": "Bone Claws", "damage": float32(5), "energy": float32(1.15), "passive": false},
			map[string]interface{}{"id": 2346, "name": "Regeneration", "damage": float32(-2), "energy": float32(0.55), "passive": true},
			map[string]interface{}{"id": 2347, "name": "Adamant skeleton", "damage": float32(-10), "energy": float32(0), "passive": true},
		},
	}
	assert.NoError(t, err)
	assert.Equal(t, want, superhero)
}

func BenchmarkHambaDecodeMap(b *testing.B) {
	schema := avro.MustParse(Schema)

	superhero := map[string]interface{}{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = avro.Unmarshal(schema, Payload, &superhero)
	}
}

func TestHambaEncodeMap(t *testing.T) {
	schema := avro.MustParse(Schema)

	superhero := map[string]interface{}{
		"id":             int32(234765),
		"affiliation_id": int32(9867),
		"name":           "Wolverine",
		"life":           float32(85.25),
		"energy":         float32(32.75),
		"powers": []interface{}{
			map[string]interface{}{"id": int32(2345), "name": "Bone Claws", "damage": float32(5), "energy": float32(1.15), "passive": false},
			map[string]interface{}{"id": int32(2346), "name": "Regeneration", "damage": float32(-2), "energy": float32(0.55), "passive": true},
			map[string]interface{}{"id": int32(2347), "name": "Adamant skeleton", "damage": float32(-10), "energy": float32(0), "passive": true},
		},
	}

	b, err := avro.Marshal(schema, &superhero)

	assert.NoError(t, err)
	assert.Equal(t, Payload, b)
}

func BenchmarkHambaEncodeMap(b *testing.B) {
	schema := avro.MustParse(Schema)

	superhero := map[string]interface{}{
		"id":             int32(234765),
		"affiliation_id": int32(9867),
		"name":           "Wolverine",
		"life":           float32(85.25),
		"energy":         float32(32.75),
		"powers": []interface{}{
			map[string]interface{}{"id": int32(2345), "name": "Bone Claws", "damage": float32(5), "energy": float32(1.15), "passive": false},
			map[string]interface{}{"id": int32(2346), "name": "Regeneration", "damage": float32(-2), "energy": float32(0.55), "passive": true},
			map[string]interface{}{"id": int32(2347), "name": "Adamant skeleton", "damage": float32(-10), "energy": float32(0), "passive": true},
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = avro.Marshal(schema, &superhero)
	}
}
