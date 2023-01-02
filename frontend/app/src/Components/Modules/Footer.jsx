import React from 'react'
import PropTypes from 'prop-types'
import { Center, Divider, Stack, Text } from '@chakra-ui/react'

export const Footer = () => (
	<footer>
		<Stack bgColor="blue.400">
			<Center>
				<Text color="snow">Footer</Text>
			</Center>
		</Stack>
	</footer>
)

Footer.propTypes = {}
