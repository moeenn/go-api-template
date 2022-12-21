package users

const GET_USER_BY_ID_QUERY = `
	SELECT * FROM users
	WHERE id = $1
`

const CREATE_USER_QUERY = `
	INSERT INTO users 
		(id, name, email, password)
	VALUES
		(:id, :name, :email, :password)
`
