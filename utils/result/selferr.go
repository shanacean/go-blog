package result

type Code int

const (
	SUCCESS Code = 200
	ERROR   Code = 500

	ERROR_USERNAME_USED    Code = 1001
	ERROR_PASSWORD_WRONG   Code = 1002
	ERROR_USER_NOT_EXIST   Code = 1003
	ERROR_TOKEN_EXIST      Code = 1004
	ERROR_TOKEN_RUNTIME    Code = 1005
	ERROR_TOKEN_WRONG      Code = 1006
	ERROR_TOKEN_TYPE_WRONG Code = 1007
	ERROR_USER_NO_RIGHT Code = 1008

	ERROR_CATEGORY_NAME_USED Code = 2001
	ERROR_CATEGORY_NOT_EXIST Code = 2002

	ERROR_ARTICLE_NOT_EXIST Code = 3001
)

var codeMap = map[Code]string{
	SUCCESS: "OK",
	ERROR: "FAIL",
	ERROR_USERNAME_USED: "Username existed",
	ERROR_PASSWORD_WRONG: "Password wrong",
	ERROR_USER_NOT_EXIST: "User doesn't exist",
	ERROR_TOKEN_EXIST: "Token doesn't exist",
	ERROR_TOKEN_RUNTIME: "Token has expired",
	ERROR_TOKEN_WRONG: "Token wrong",
	ERROR_TOKEN_TYPE_WRONG: "Token Type wrong",
	ERROR_USER_NO_RIGHT: "NO RIGHT",

	ERROR_CATEGORY_NAME_USED: "Category Name existed",
	ERROR_CATEGORY_NOT_EXIST: "Category doesn't exist",

	ERROR_ARTICLE_NOT_EXIST: "Article doesn't exist",
}

func (c Code) GetCodeMessage() string {
	return codeMap[c]
}