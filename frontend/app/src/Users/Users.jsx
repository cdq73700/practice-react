import { Stack } from '@chakra-ui/react'
import { skipToken } from '@reduxjs/toolkit/dist/query'
import React from 'react'
import { useParams } from 'react-router-dom'
import { useUserListQuery } from '../Api'
import { ItemList } from './ItemList'

export const Users = () => {
	const { id } = useParams()
	const userList = useUserListQuery()

	return (
		<>
			<ItemList {...userList} />
		</>
	)
}
Users.propTypes = {}
