import React from 'react'
import PropTypes from 'prop-types'
import { Center, Stack, Text } from '@chakra-ui/react'

export const Header = () => (
	<header>
		<Stack bgColor="blue.400">
			<Center>
				<Text color="snow">Header</Text>
			</Center>
		</Stack>
	</header>
)

Header.propTypes = {}
