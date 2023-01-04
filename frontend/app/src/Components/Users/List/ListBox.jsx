import { Card, CardBody, Stack, Text } from '@chakra-ui/react'
import PropTypes from 'prop-types'
import React from 'react'

export const ListBox = ({ items }) => (
	<Stack p="5px 15px 0px 15px" h="100vh" overflowY="scroll">
		{items.map((item) => (
			<Card key={item.Id} variant="outline">
				<CardBody>
					<Text>{item.Email}</Text>
				</CardBody>
			</Card>
		))}
	</Stack>
)

ListBox.propTypes = {
	items: PropTypes.array,
}
