import {
	Button,
	ButtonGroup,
	Card,
	CardBody,
	CardFooter,
	CardHeader,
	Heading,
	HStack,
	Spacer,
	Spinner,
	Stack,
	Text,
	useDisclosure,
} from '@chakra-ui/react'
import { AnimatePresence } from 'framer-motion'
import PropTypes from 'prop-types'
import React, { useCallback, useEffect, useRef } from 'react'
import { AnimationButtonFade } from '../../Button/Button'
import { useSelectUser } from '../Reducers/CustomHook'
import { Edit } from './Edit'
import { View } from './View'

export const UserItem = ({
	isUninitialized,
	isLoading,
	isFetching,
	isError,
	error,
	data,
}) => {
	const { userId } = useSelectUser()
	const { isOpen, onOpen, onClose } = useDisclosure()
	const EditOnClick = useCallback(() => {
		if (!isOpen) {
			onOpen()
		} else {
			onClose()
		}
	}, [isOpen])

	if (isUninitialized) {
		return (
			<>
				<Text>isUninitialized</Text>
			</>
		)
	}
	if (isError) {
		console.log(error)
		return (
			<>
				<Text>Error</Text>
			</>
		)
	}

	useEffect(() => {
		onClose()
	}, [userId])

	return isLoading ? (
		<Spinner />
	) : (
		data && (
			<Stack p="5px 15px 0px 15px" maxH="100vh">
				<Card variant="outline">
					<CardHeader>
						<HStack>
							<Heading>{data.Name}</Heading>
							<Spacer />
							<ButtonGroup>
								<Button
									colorScheme="green"
									onClick={EditOnClick}>
									Edit
								</Button>
							</ButtonGroup>
						</HStack>
					</CardHeader>
					<CardBody>
						{isOpen ? (
							<Edit Email={data.Email} Name={data.Name} />
						) : (
							<View Email={data.Email} Name={data.Name} />
						)}
					</CardBody>
					<CardFooter>
						<Spacer />
						<AnimatePresence>
							{isOpen && (
								<ButtonGroup>
									<AnimationButtonFade>
										Save
									</AnimationButtonFade>
									<AnimationButtonFade colorScheme="red">
										Delete
									</AnimationButtonFade>
								</ButtonGroup>
							)}
						</AnimatePresence>
					</CardFooter>
				</Card>
			</Stack>
		)
	)
}

UserItem.propTypes = {
	isUninitialized: PropTypes.bool.isRequired,
	isLoading: PropTypes.bool.isRequired,
	isFetching: PropTypes.bool.isRequired,
	isError: PropTypes.bool.isRequired,
	error: PropTypes.any,
	data: PropTypes.any,
}
