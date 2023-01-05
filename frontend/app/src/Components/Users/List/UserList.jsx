import { Spinner, Stack, Text } from '@chakra-ui/react'
import PropTypes from 'prop-types'
import React, { useEffect } from 'react'
import { useSelectUser } from '../Reducers/CustomHook'
import { ListBox } from './ListBox'

export const UserList = ({
	isUninitialized,
	isLoading,
	isFetching,
	isError,
	error,
	data,
}) => {
	const { userId, updateUser } = useSelectUser()
	if (isUninitialized) {
		return (
			<>
				<Text>isUninitialized</Text>
			</>
		)
	}
	if (isError) {
		console.log(error)
		return (
			<>
				<Text>Error</Text>
			</>
		)
	}

	useEffect(() => {
		if (data && userId === '') {
			updateUser(data.items[0].Id)
		}
	}, [data])

	return <>{isLoading ? <Spinner /> : data && <ListBox {...data} />}</>
}

UserList.propTypes = {
	isUninitialized: PropTypes.bool.isRequired,
	isLoading: PropTypes.bool.isRequired,
	isFetching: PropTypes.bool.isRequired,
	isError: PropTypes.bool.isRequired,
	error: PropTypes.any,
	data: PropTypes.any,
}
