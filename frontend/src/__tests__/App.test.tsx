import { render, screen } from '@testing-library/react'
import { MemoryRouter } from 'react-router-dom'
import App from '../App'

jest.mock('../components/NoiseOverlay', () => () => <div />)
jest.mock('../components/StarField', () => () => <div />)

it('renders dashboard heading', () => {
  render(
    <MemoryRouter>
      <App />
    </MemoryRouter>
  )
  expect(screen.getByRole('heading', { name: /web crawler/i })).toBeTruthy()
})
