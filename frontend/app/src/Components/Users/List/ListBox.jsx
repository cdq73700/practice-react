import { Card, CardBody, Stack, Text } from '@chakra-ui/react'
import { skipToken } from '@reduxjs/toolkit/dist/query'
import PropTypes from 'prop-types'
import React, { useCallback } from 'react'
import { useLazyUserDetailQuery, usePrefetch } from '../Api'
import { useSelectUser } from '../Reducers/CustomHook'

export const ListBox = ({ items }) => {
	const { updateUser } = useSelectUser()
	const prefetch = usePrefetch('userDetail')

	const CardOnClick = useCallback((Id) => {
		updateUser(Id)
	}, [])
	return (
		<Stack p="5px 15px 0px 15px" h="100vh" overflowY="scroll">
			{items.map((item) => (
				<Card key={item.Id} variant="outline" cursor="pointer">
					<CardBody
						onMouseOver={() => {
							prefetch(item.Id)
						}}
						onClick={() => CardOnClick(item.Id)}>
						<Text>{item.Email}</Text>
					</CardBody>
				</Card>
			))}
		</Stack>
	)
}

ListBox.propTypes = {
	items: PropTypes.array,
}
