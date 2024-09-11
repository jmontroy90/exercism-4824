package day1

// So the thing about this problem is that I don't really see how people get to the standard boolean chain method
// through just parsing a description of leap years. Natural language leads much more to a nested-if branching
// solution, even though they're totally equivalent. I think more comfort with equivalencies between if-else
// statements vs. boolean chains will go a long way in helping me parse these problems.

// There's also the interesting angle via truth tables -- the standard solution relies on certain combinations
// of boolean inputs between our three variables being impossible; namely, if it's divisible by 400, it's divisible
// by 4 and 100 too, because math. See #5 for more details of how I explored that problem. Would probably need
// guidance from someone much more trained at this than me.

// Overall, I find #1 to be the cleanest-looking, #4 (early-exit) to be the easiest to think about,
// #5 to be the most brainless / methodical one (but also totally impossible to understand intuitively),
// and #2 (nested-if) to be the most literal. When the problem is written like #3, I find that easy to
// parse and code. Most people choose #1 and #2 as solutions.

// classic "boolean chain" method
// my brain doesn't "leap" (HAHA) to this solution yet, but the general logic is understandable.
// either it's gotta be divisible by 4 and not 100 and not 400, OR 4 AND 100 AND 400
func IsLeapYear(year int) bool {
	by4 := year%4 == 0
	by100 := year%100 == 0
	by400 := year%400 == 0
	// a && (!b OR C)
	return by4 && (!by100 || by400)
}

// "branching" / "nested-if" method; this solution is I think the most literal interpretation of the instructions.
// In every year that is evenly divisible by 4.
// Unless the year is evenly divisible by 100,
// in which case it's only a leap year if the year is also evenly divisible by 400.
func IsLeapYear2(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			return year%400 == 0
		} else {
			return true
		}
	} else {
		return false
	}
}

// "rewrite cases" method
// * Evenly divisible by 4, unless the year is evenly divisible by 100.
// * Evenly divisible by 400.
func IsLeapYear3(year int) bool {
	if year%4 == 0 {
		return year%100 != 0
	}
	if year%400 == 0 {
		return true
	}
	return false
}

// "rule out cases" method
// tbh for readability i prefer this; it is basically an early-exit / short-circuit equivalent
func IsLeapYear4(year int) bool {
	if year%4 != 0 {
		return false
	}
	if year%100 == 0 {
		return year%400 == 0
	}
	return true
}

// "truth table" method; derived by truth table and OR, maybe some way to simplify this expression.
// this is the most reliable way FOR ME to figure out the right boolean chain, but -
// it's definitely the least readable and understandable.
//
// It's worth noting that this boolean expression and the one in #1 do NOT produce the same outputs!
// Like, if somehow the year was divisible by 100 and not divisible by 4, this expression would
// still produce false - #1's expression would say true. The only reason that isn't a problem for #1 is that
// those inputs are impossible because of how the boolean inputs are correlated, being all modulo subcases
// of one another. I'm not really sure how to think about that - I think I'm not in a crazy dumb place,
// but even if I am, I think it's a necessary crazy dumb place. I think if these boolean inputs were
// independent, this would be the right answer for the outcomes we want.
func IsLeapYear5(year int) bool {
	by4 := year%4 == 0
	by100 := year%100 == 0
	by400 := year%400 == 0
	// a && (b XNOR C)
	//return by4 && (by100 == by400)
	return by4 && (!by100 && !by400 || by100 && by400)
}

/*
Here's yer truth table --

by4		by100		by400		result		equation
0 		0 			0 			0
1 		1 			0 			0
1 		0 			0 			1  			by4 && !by100 && !by400
1 		1 			1 			1 			by4 && by100 && by400

1 		0 			1 			impossible
0 		1 			1 			impossible
0 		0 			1 			impossible
0 		1 			0 			impossible

= (by4 && !by100 && !by400) || (by4 && by100 && by400)
	= by4 && (!by100 && !by400 || by100 && by400)
*/

// Example: "A person is Good if they are rad (a), unless they are also sad (b), in which case
// they are only Good if they're also sick (c)".

/*
rad		sad		sick		result		equation
0 		0 			0 			0
1 		1 			0 			0
1 		0 			0 			1  		a && !b && !c
1 		1 			1 			1		a && b && c
1 		0 			1 			1		a && !b && c
0 		1 			1 			0
0 		0 			1 			0
0 		1 			0 			0

= (a && !b && !c) || (a && b && c) || (a && !b && c)
	= a && (!b || c) (I used Wolfram Alpha, I'm no good at boolean simplification)
	https://www.wolframalpha.com/input?i=simplify+%28a+%26%26+%21b+%26%26+%21c%29+%7C%7C+%28a+%26%26+b+%26%26+c%29+%7C%7C+%28a+%26%26+%21b+%26%26+c%29

Well well well, look who it is. The truth table expression for three independent inputs is exactly #1's standard solution.
I wonder if this is secretly how we got it. But see, it only works because "rad, not sad, and sick" is impossible
in our setup, because it'd be "divisible by 4, not by 100, but yes also by 400", which is no number.

So what's the verdict? I still find the "standard" answer not very intuitive, but for independent boolean inputs,
you get that same boolean expression -- and it's not so bad to think about. For correlated inputs (what we have), you get the weird XNOR one I did. Since
so many input combos are impossible for these booleans, I guess the choice becomes between the technically-correct
and the more-interpretable -- I would go interpretable every time.
*/
