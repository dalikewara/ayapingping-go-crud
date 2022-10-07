# Changelogs

### 2022

- v1.7.0 (2022-10-04)
  - add `ByteCondition` properties: `AtLeastHaveUpperLetterCount`, `AtLeastHaveLowerLetterCount`, `AtLeastHaveNumberCount`, `AtLeastHaveSpecialCharCount`
  - add some char bytes

- v1.6.0 - v1.6.2 (2022-09-27)
  - optimize `Byte` validate logic
  - change function `Bytes` to `Byte` and change its first param from `[]byte` to `string`
  - this function now save to validate only ASCII characters

- v1.4.0 - v1.5.1 (2022-09-26)
  - add proper benchmark
  - seperate validate function to `Bytes` and `String`
  - optimize validate logic

- v1.3.0 (2022-09-25)
  - change function form 
  - optimize validate logic

- v1.2.2 (2022-09-23)
  - optimize validate logic

- v1.1.2 - v1.2.1 (2022-09-22)
  - optimize validate logic
  - Fix possibly issue on `setWords`

- v1.0.1 - v1.1.1 (2022-09-21)
  - optimize validate logic
  - add methods: `OnlyContainPrefixWords`, `OnlyContainSuffixWords`, `MustNotContainPrefixWords`, `MustNotContainSuffixWords`
  - add constants: `SPECIALCHARS`, `QUOTES`, `BRACKETS`, `OPERATORS`

- v1.0.0 (2022-09-20)
    - Initial release
