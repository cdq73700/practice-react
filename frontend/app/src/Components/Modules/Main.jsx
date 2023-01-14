import { Grid, GridItem, SimpleGrid, Stack } from '@chakra-ui/react'
import PropTypes from 'prop-types'
import React from 'react'

export const Main = ({ children }) => (
	<Grid gridTemplateColumns="150px 1fr 150px">
		<GridItem>
			<Stack bgColor="gray.200" h="100%" />
		</GridItem>
		<GridItem>
			<Stack>{children}</Stack>
		</GridItem>
		<GridItem>
			<Stack bgColor="gray.200" h="100%" />
		</GridItem>
	</Grid>
)

Main.propTypes = {
	children: PropTypes.node.isRequired,
}
