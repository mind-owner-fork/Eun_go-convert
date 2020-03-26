package convert

import (
	"strconv"
	"time"
)

func (stdRecipes) intToInt64(c Converter, in int, out *int64) error {
	*out = int64(in)
	return nil
}
func (stdRecipes) int8ToInt64(c Converter, in int8, out *int64) error {
	*out = int64(in)
	return nil
}
func (stdRecipes) int16ToInt64(c Converter, in int16, out *int64) error {
	*out = int64(in)
	return nil
}
func (stdRecipes) int32ToInt64(c Converter, in int32, out *int64) error {
	*out = int64(in)
	return nil
}
func (stdRecipes) int64ToInt64(c Converter, in int64, out *int64) error {
	*out = in
	return nil
}
func (stdRecipes) uintToInt64(c Converter, in uint, out *int64) error {
	*out = int64(in)
	return nil
}
func (stdRecipes) uint8ToInt64(c Converter, in uint8, out *int64) error {
	*out = int64(in)
	return nil
}
func (stdRecipes) uint16ToInt64(c Converter, in uint16, out *int64) error {
	*out = int64(in)
	return nil
}
func (stdRecipes) uint32ToInt64(c Converter, in uint32, out *int64) error {
	*out = int64(in)
	return nil
}
func (stdRecipes) uint64ToInt64(c Converter, in uint64, out *int64) error {
	*out = int64(in)
	return nil
}
func (stdRecipes) boolToInt64(c Converter, in bool, out *int64) error {
	switch in {
	case true:
		*out = 1
	default:
		*out = 0
	}
	return nil
}

func (stdRecipes) float32ToInt64(c Converter, in float32, out *int64) error {
	*out = int64(in)
	return nil
}
func (stdRecipes) float64ToInt64(c Converter, in float64, out *int64) error {
	*out = int64(in)
	return nil
}
func (stdRecipes) stringToInt64(c Converter, in string, out *int64) error {
	if in == "" {
		*out = 0
		return nil
	}
	i, err := strconv.ParseInt(in, 0, 64)
	if err != nil {
		return err
	}
	*out = i
	return nil
}
func (stdRecipes) timeToInt64(c Converter, in time.Time, out *int64) error {
	*out = in.Unix()
	return nil
}
