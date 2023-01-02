import { Spacer, Stack } from '@chakra-ui/react'
import React from 'react'
import Div100vh from 'react-div-100vh'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { App } from '../App'
import { Footer } from '../Components/Modules/Footer'
import { Header } from '../Components/Modules/Header'
import { Users } from '../Users/Users'

export const RouterConfig = () => (
	<>
		<BrowserRouter>
			<Div100vh>
				<Header />
				<body>
					<Routes>
						<Route path="/" element={<App />} />
						<Route path="about" element={<div>About</div>} />
						<Route path="users" element={<Users />}>
							<Route path=":id" element={<Users />} />
						</Route>
					</Routes>
				</body>
				<Footer />
			</Div100vh>
		</BrowserRouter>
	</>
)
