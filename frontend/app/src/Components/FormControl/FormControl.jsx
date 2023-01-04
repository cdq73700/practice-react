import React from 'react'
import PropTypes from 'prop-types'
import {
	Center,
	Flex,
	FormControl as ChakraFormControl,
	FormLabel,
	Grid,
	GridItem,
	Input,
	Text,
} from '@chakra-ui/react'

export const FormControlView = ({ Title, Value }) => (
	<ChakraFormControl>
		<Grid gridTemplateColumns="0.15fr 1fr" h="50px">
			<GridItem>
				<FormLabel h="100%">
					<Center h="100%">{Title}</Center>
				</FormLabel>
			</GridItem>
			<GridItem>
				<Flex align="center" h="100%">
					<Text>{Value}</Text>
				</Flex>
			</GridItem>
		</Grid>
	</ChakraFormControl>
)

FormControlView.propTypes = {
	Title: PropTypes.string.isRequired,
	Value: PropTypes.string.isRequired,
}

export const FormControlEdit = ({ Title, Value, InputType }) => (
	<ChakraFormControl>
		<Grid gridTemplateColumns="0.15fr 1fr" h="50px">
			<GridItem>
				<FormLabel h="100%">
					<Center h="100%">{Title}</Center>
				</FormLabel>
			</GridItem>
			<GridItem>
				<Flex align="center" h="100%">
					<Input type={InputType} defaultValue={Value} />
				</Flex>
			</GridItem>
		</Grid>
	</ChakraFormControl>
)

FormControlEdit.propTypes = {
	Title: PropTypes.string.isRequired,
	Value: PropTypes.string.isRequired,
	InputType: PropTypes.string.isRequired,
}
