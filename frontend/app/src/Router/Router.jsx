import React from 'react'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { App } from '../App'
import { Users } from '../Users/Users'

export const RouterConfig = () => (
	<>
		<BrowserRouter>
			<Routes>
				<Route path="/" element={<App />} />
				<Route path="about" element={<div>About</div>} />
				<Route path="users" element={<Users />}>
					<Route path=":id" element={<Users />} />
				</Route>
			</Routes>
		</BrowserRouter>
	</>
)
