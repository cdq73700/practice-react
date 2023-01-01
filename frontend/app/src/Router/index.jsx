import React from 'react'
import { RouterProvider as ReactRouterProvider } from 'react-router-dom'
import { Provider } from 'react-redux'
import { ThemeProvider, ColorModeProvider } from '@chakra-ui/react'
import PropTypes from 'prop-types'
import { router } from './Router'
import { store } from '../Store'

/**
 * ルートプロバイダー
 * @param {node} children
 * @returns
 */
export const RouterProvider = ({ children }) => (
	<Provider store={store}>
		<ReactRouterProvider router={router}>
			<ThemeProvider>
				<ColorModeProvider>{children}</ColorModeProvider>
			</ThemeProvider>
		</ReactRouterProvider>
	</Provider>
)

RouterProvider.propTypes = {
	children: PropTypes.node,
}
