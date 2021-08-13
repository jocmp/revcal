# Revcal

This is a Go rewrite of a Ruby library that started with
the work of Jonathan Badger ([GitHub source](https://github.com/jhbadger/FrenchRevCal-ruby)). The
original source code comments, Usage and Algorithm documentation are only slightly modified
to conform to Go verbage.

## Usage

The script "revcal" without any arguments returns the current date in
the revolutionary calendar, along with the associated concept for the
day. Providing a date (in the standard ISO 8601 date format that the time package
can parse), returns the value for that date instead (which is useful
for looking up one's birthday, for instance.)

```bash
 ./revcal 2021-01-02
Today is 13 Nivôse 229 celebrating the Ardoise (Slate).
./revcal
Today is 26 Thermidor 229 celebrating the Myrte (Myrtle).
```

revcal is a script that converts standard Gregorian dates
into those used by the [French Republican Calendar](http://en.wikipedia.org/wiki/French_Republican_Calendar).
This was the calendar used by the Republican government in France between 1793 and 1805.
Its structure reflects the secular, rational ideals of the time,
with 12 months of a uniform 30 days, broken into 3 10-day "weeks". The
remaining five days of the year (or 6 in leap years) were filled with
monthless year-end celebrations. Year 1 of the calendar started in
September 1792, with the official founding of the French
Republic. Each day of the year was associated with a plant, animal, or
tool, replacing the saints of Catholic tradition.

## Algorithm

The arithmetic of the date package is used to calculate the number of days
since September 22, 1792 (Day 1 of Year 1). Then it is a simple matter
of jumping forward by the number of days in each year until the number
of remaining days is less than a length of a year. Then dividing this
number by 30 gives the month, and the modulo the day. The only real
issue is what counts as a leap year. There are several conflicting
ideas as to the method of computing the revolutionary leap years.
The chosen method is to treat the historical leap years as 3, 7, 11, 15, and 20,
and to treat subsequent years divisible by 4 that aren't divisible by 100
(unless they are also divisible by 400), as per the Gregorian method.

There are several Go tests to confirm conversion of famous dates
such as 18 Brumaire 8 (Nov 27, 1799; the date of Napoleon Bonaparte's coup against the French Directory,
effectively killing the revolution as an ideological movement), the rise of the
July Monarchy in 1830, as well as the return of the Bonaparte dynasty for one last time in
1851 with the rise of Napolean III.

## Other notes

* I've restructured the Ruby version here https://github.com/jocmp/revcal_rb
* Follows the [golang-standards](https://github.com/golang-standards/project-layout) project layout.
