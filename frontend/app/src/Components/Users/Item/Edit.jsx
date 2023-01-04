import { Stack } from '@chakra-ui/react'
import PropTypes from 'prop-types'
import React from 'react'
import { FormControlEdit } from '../../FormControl/FormControl'

export const Edit = ({ Email, Name }) => (
	<Stack gap={3}>
		<FormControlEdit Title="Email" Value={Email} InputType="email" />
		<FormControlEdit Title="Name" Value={Name} InputType="text" />
	</Stack>
)

Edit.propTypes = {
	Email: PropTypes.string,
	Name: PropTypes.string,
}
