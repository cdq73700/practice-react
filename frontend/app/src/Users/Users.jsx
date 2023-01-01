import { Stack } from '@chakra-ui/react'
import React from 'react'
import { useUserListQuery } from '../Api'
import { ItemList } from './ItemList'

export const Users = () => {
	const users = Array.from({ length: 10 }, (_, index) => {
		const id = index + 1
		return {
			id,
			email: `test${id}@example.com`,
			username: `test${id}`,
			password: `test${id}`,
		}
	})

	const userList = useUserListQuery()

	console.log(userList)

	return (
		<Stack>
			<ItemList {...userList} />
		</Stack>
	)
}
Users.propTypes = {}
