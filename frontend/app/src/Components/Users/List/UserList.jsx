import { Spinner, Stack, Text } from '@chakra-ui/react'
import PropTypes from 'prop-types'
import React from 'react'
import { ListBox } from './ListBox'

export const UserList = ({
	isUninitialized,
	isLoading,
	isFetching,
	isError,
	error,
	data,
}) => {
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
