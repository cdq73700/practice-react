import { Grid, GridItem, SimpleGrid, Stack } from '@chakra-ui/react'
import PropTypes from 'prop-types'
import React from 'react'

export const Main = ({ children }) => (
	<Grid gridTemplateColumns="150px 1fr 150px" minH="100vh" maxH="100vh">
		<GridItem>
			<Stack bgColor="gray.200" h="100%" />
		</GridItem>
		<GridItem>
			<Stack maxH="100vh">{children}</Stack>
		</GridItem>
		<GridItem>
			<Stack bgColor="gray.200" h="100%" />
		</GridItem>
	</Grid>
)

Main.propTypes = {
	children: PropTypes.node.isRequired,
}
