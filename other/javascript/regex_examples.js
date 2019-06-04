var testString = 'He11ooo thereeee regexFindRepeatedNeighboringCharacterswhy ar3 you her3?';

// \d is digit
// + is more than Matches the preceding expression 1 or more times. Equivalent to {1,}.
var regexFindAllDigits = /\d+/g;
console.log(testString.match(regexFindAllDigits));

// The backreference \1 (backslash one) references the first capturing group. 
// \1 matches the exact same text that was matched by the first capturing group.
// https://www.regular-expressions.info/backref.html
var regexRepeatedCharacters = /(.)\1+/g;
console.log(testString.match(regexRepeatedCharacters));

// % operator = Modulus (division remainder)	
// 7 % 2 = 1
// 7 / 2 = 3 with remainder of 1
var int = 8;

// The conditional operator assigns a value to a variable based on a condition.
// var = (int % 2) ? 0 : 9	voteable = (age < 18) ? "Too young":"Old enough";
console.log(int % 2 ? 0 : 9, int - 1)