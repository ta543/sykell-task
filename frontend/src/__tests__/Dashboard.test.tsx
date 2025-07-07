import { render, screen, waitFor, fireEvent } from '@testing-library/react'
import Dashboard from '../pages/Dashboard'
import { MemoryRouter } from 'react-router-dom'
import axios from 'axios'

jest.mock('axios')
const mockedAxios = axios as jest.Mocked<typeof axios>

describe('Dashboard', () => {
  it('fetches and displays urls', async () => {
    mockedAxios.get.mockResolvedValueOnce({ data: [
      { id: 1, title: 'Example', address: 'http://ex.com', html_version: 'HTML5', internal_links: 1, external_links: 2, broken_links: 0, status: 'done' }
    ] })
    render(<MemoryRouter><Dashboard /></MemoryRouter>)
    await waitFor(() => expect(mockedAxios.get).toHaveBeenCalled())
    expect(await screen.findByText('Example')).toBeTruthy()
  })

  it('adds new url', async () => {
    mockedAxios.get.mockResolvedValue({ data: [] })
    mockedAxios.post.mockResolvedValue({})
    render(<MemoryRouter><Dashboard /></MemoryRouter>)
    fireEvent.change(screen.getByPlaceholderText(/enter url/i), { target: { value: 'http://new.com' } })
    fireEvent.click(screen.getByText('Add'))
    await waitFor(() => expect(mockedAxios.post).toHaveBeenCalledWith('/api/urls', { address: 'http://new.com' }, { headers: { Authorization: 'secret-token' } }))
  })
})
