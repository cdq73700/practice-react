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
			<Stack>
				<Text>isUninitialized</Text>
			</Stack>
		)
	}
	if (isError) {
		return (
			<Stack>
				<Text>Error</Text>
			</Stack>
		)
	}

	return (
		<Stack>
			{isLoading ? (
				<Spinner />
			) : (
				data &&
				data.items.map((items) => (
					<Stack key={items.Email}>
						<Item isFetching={isFetching} data={items} />
					</Stack>
				))
			)}
		</Stack>
	)
}

ItemList.propTypes = {
	isUninitialized: PropTypes.bool.isRequired,
	isLoading: PropTypes.bool.isRequired,
	isFetching: PropTypes.bool.isRequired,
	isError: PropTypes.bool.isRequired,
	data: PropTypes.any,
}
