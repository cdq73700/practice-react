import { ChakraBaseProvider, theme } from '@chakra-ui/react'
import React from 'react'
import { Provider } from 'react-redux'
import { store } from '../Store'
import { RouterConfig } from './Router'

/**
 * ルートプロバイダー
 * @param {node} children
 * @returns
 */
export const RouterProvider = () => (
	<Provider store={store}>
		<ChakraBaseProvider theme={theme}>
			<RouterConfig />
		</ChakraBaseProvider>
	</Provider>
)

RouterProvider.propTypes = {}
