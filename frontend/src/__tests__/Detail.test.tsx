import { render, screen, waitFor } from '@testing-library/react'
import Detail from '../pages/Detail'
import { MemoryRouter, Route, Routes } from 'react-router-dom'
import axios from 'axios'

jest.mock('axios')
const mockedAxios = axios as jest.Mocked<typeof axios>

describe('Detail', () => {
  it('loads url detail', async () => {
    mockedAxios.get.mockResolvedValueOnce({ data: {
      id: 1,
      title: 'Example',
      address: 'http://ex.com',
      html_version: 'HTML5',
      internal_links: 1,
      external_links: 2,
      broken_links: 0,
      broken_links_detail: [],
      status: 'done',
      h1_count: 1,
      h2_count: 0,
      h3_count: 0,
      has_login_form: false,
    }})
    render(
      <MemoryRouter initialEntries={["/detail/1"]}>
        <Routes>
          <Route path="/detail/:id" element={<Detail />} />
        </Routes>
      </MemoryRouter>
    )
    await waitFor(() => expect(mockedAxios.get).toHaveBeenCalled())
    expect(await screen.findByText('Example')).toBeTruthy()
    expect(screen.getByText('HTML Version: HTML5')).toBeTruthy()
  })
})
