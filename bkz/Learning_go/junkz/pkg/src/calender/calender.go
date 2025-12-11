package calender

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	year  int
	month int
	day   int
}

const (
	Jan = iota + 1
	Feb
	Mar
	Apr
	May
	Jun
	Jul
	Aug
	Sep
	Oct
	Nov
	Dec
)

func NewDate(year, month, day int) *Date {
	date := &Date{
		year:  year,
		month: month,
		day:   day,
	}
	yerr := date.SetYear(year)
	if yerr != nil {
		log.Fatal(yerr)
	}

	merr := date.SetMonth(month)
	if merr != nil {
		log.Fatal(merr)
	}

	derr := date.SetDay(day)
	if derr != nil {
		log.Fatal(derr)
	}

	return date
}

func (d *Date) SetYear(year int) error {
	if year < 1100 {
		return errors.New("invalid year value")
	}
	d.year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month value")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	switch d.month {
	case Apr, Jun, Sep, Nov:
		if day < 1 || day > 30 {
			return errors.New("invalid day value")
		}
	case Jan, Mar, May, Jul, Aug, Oct, Dec:
		if day < 1 || day > 31 {
			return errors.New("invalid day value")
		}
	case Feb:
		leapYear := (d.year%4 == 0) && (d.year%400 != 0)
		if leapYear {
			if day < 1 || day > 29 {
				return errors.New("invalid day value")
			}
		} else {
			if day == 29 {
				return errors.New("Not a leap year.")
			}
			if day < 1 || day > 28 {
				return errors.New("invalid day value")
			}
		}
	default:
		return errors.New("wrong case!")
	}
	d.day = day
	return nil
}

func (d *Date) Stringer() string {
	return fmt.Sprintf("Date: %d/%d/%d\n", d.month, d.day, d.year)
}
