import { Stack } from '@chakra-ui/react'
import { createContext, useReducer } from 'react'
import { UsersView } from '../View/View'

const initialState = {
	userId: '',
}

export const UserContext = createContext()
export const UserDispatch = createContext()

const reducer = (state, action) => {
	switch (action.type) {
		case 'SELECT_USER':
			return { userId: action.payload }

		default:
			break
	}
	return state
}

export const Users = () => {
	const [state, dispatch] = useReducer(reducer, initialState)

	return (
		<UserContext.Provider value={state}>
			<UserDispatch.Provider value={dispatch}>
				<Stack>
					<UsersView />
				</Stack>
			</UserDispatch.Provider>
		</UserContext.Provider>
	)
}
