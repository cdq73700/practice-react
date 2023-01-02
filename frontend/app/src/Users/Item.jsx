import React from 'react'
import PropTypes from 'prop-types'
import { HStack, Skeleton, Stack, Text } from '@chakra-ui/react'

const EmailFiled = ({ Email }) => (
	<HStack>
		<Text>Email:</Text>
		<Text>{Email}</Text>
	</HStack>
)

EmailFiled.propTypes = {
	Email: PropTypes.string,
}

const UsernameFiled = ({ Username }) => (
	<HStack>
		<Text>Username:</Text>
		<Text>{Username}</Text>
	</HStack>
)

UsernameFiled.propTypes = {
	Username: PropTypes.string,
}

export const Item = ({ isFetching, data }) => (
	<Skeleton isLoaded={!isFetching}>
		<EmailFiled {...data} />
		<UsernameFiled {...data} />
	</Skeleton>
)

Item.propTypes = {
	isFetching: PropTypes.bool.isRequired,
	data: PropTypes.any,
}
