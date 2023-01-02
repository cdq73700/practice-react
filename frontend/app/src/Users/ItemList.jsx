import React from 'react'
import PropTypes from 'prop-types'
import { Spinner, Stack, Text } from '@chakra-ui/react'
import { Item } from './Item'

export const ItemList = ({
	isUninitialized,
	isLoading,
	isFetching,
	isError,
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
		return (
			<>
				<Text>Error</Text>
			</>
		)
	}

	return (
		<>
			{isLoading ? (
				<Spinner />
			) : (
				data &&
				data.items.map((items) => (
					<Stack key={items.Id}>
						<Item isFetching={isFetching} data={items} />
					</Stack>
				))
			)}
		</>
	)
}

ItemList.propTypes = {
	isUninitialized: PropTypes.bool.isRequired,
	isLoading: PropTypes.bool.isRequired,
	isFetching: PropTypes.bool.isRequired,
	isError: PropTypes.bool.isRequired,
	data: PropTypes.any,
}
