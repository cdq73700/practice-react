import React from 'react'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { App } from '../App'
import { Footer } from '../Components/Modules/Footer'
import { Header } from '../Components/Modules/Header'
import { Main } from '../Components/Modules/Main'
import { Users } from '../Components/Users/Reducers/Reducers'

export const RouterConfig = () => (
	<>
		<BrowserRouter>
			<Header />
			<Main>
				<Routes>
					<Route path="/" element={<App />} />
					<Route path="about" element={<div>About</div>} />
					<Route path="users" element={<Users />} />
				</Routes>
			</Main>
			<Footer />
		</BrowserRouter>
	</>
)
