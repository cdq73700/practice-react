import { ColorModeProvider, ThemeProvider } from '@chakra-ui/react'
import PropTypes from 'prop-types'
import React from 'react'
import { Provider } from 'react-redux'
import { store } from '../Store'
import { RouterConfig } from './Router'

/**
 * ルートプロバイダー
 * @param {node} children
 * @returns
 */
export const RouterProvider = ({ children }) => (
	<Provider store={store}>
		<RouterConfig>
			<ThemeProvider>
				<ColorModeProvider>{children}</ColorModeProvider>
			</ThemeProvider>
		</RouterConfig>
	</Provider>
)

RouterProvider.propTypes = {
	children: PropTypes.node,
}
