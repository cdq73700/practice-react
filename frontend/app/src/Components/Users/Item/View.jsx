import { Stack } from '@chakra-ui/react'
import PropTypes from 'prop-types'
import React from 'react'
import { FormControlView } from '../../FormControl/FormControl'

export const View = ({ Email, Name }) => (
	<Stack gap={3}>
		<FormControlView Title="Email" Value={Email} />
		<FormControlView Title="Name" Value={Name} />
	</Stack>
)

View.propTypes = {
	Email: PropTypes.string,
	Name: PropTypes.string,
}
