import React from 'react'
import { createBrowserRouter } from 'react-router-dom'
import { App } from '../App'
import { Users } from '../Users/Users'

export const router = createBrowserRouter([
	{
		path: '/',
		element: <App />,
	},
	{
		path: 'about',
		element: <div>About</div>,
	},
	{
		path: 'users',
		element: <Users />,
	},
])
