import React from 'react'
import PropTypes from 'prop-types'
import { Button as ChakraButton } from '@chakra-ui/react'
import { motion } from 'framer-motion'

export const AnimationButtonFade = ({ children, ...otherProps }) => (
	<ChakraButton
		as={motion.button}
		colorScheme="blue"
		initial={{ opacity: 0 }}
		animate={{ opacity: 1 }}
		exit={{ opacity: 0 }}
		transition={{ duration: 0.5 }}
		{...otherProps}>
		{children}
	</ChakraButton>
)

AnimationButtonFade.propTypes = {
	children: PropTypes.node.isRequired,
}
