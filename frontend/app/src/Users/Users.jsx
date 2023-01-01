import { Stack } from '@chakra-ui/react'
import React from 'react'
import { useUserListQuery } from '../Api'
import { ItemList } from './ItemList'

export const Users = () => {
	const userList = useUserListQuery()

	return (
		<Stack>
			<ItemList {...userList} />
		</Stack>
	)
}
Users.propTypes = {}
