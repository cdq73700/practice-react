import React from 'react'
import PropTypes from 'prop-types'
import { HStack, Skeleton, Stack, Text } from '@chakra-ui/react'

const EmailFiled = ({ Email }) => (
	<>
		<Text>{`Email: ${Email}`}</Text>
	</>
)

EmailFiled.propTypes = {
	Email: PropTypes.string,
}

const NameFiled = ({ Name }) => (
	<>
		<Text>{`Name: ${Name}`}</Text>
	</>
)

NameFiled.propTypes = {
	Name: PropTypes.string,
}

export const Item = ({ isFetching, data }) => {
	const { Email, Name } = data
	return (
		<Skeleton isLoaded={!isFetching}>
			<EmailFiled Email={Email} />
			<NameFiled Name={Name} />
		</Skeleton>
	)
}

Item.propTypes = {
	isFetching: PropTypes.bool.isRequired,
	data: PropTypes.any,
}
