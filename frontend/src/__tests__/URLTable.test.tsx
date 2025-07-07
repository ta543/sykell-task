import { render, screen, fireEvent } from '@testing-library/react'
import { MemoryRouter } from 'react-router-dom'
import URLTable from '../components/URLTable'

const items = [
  { id: 1, title: 'Example', address: 'http://ex.com', html_version: 'HTML5', internal_links: 1, external_links: 2, broken_links: 0, status: 'done' }
]

describe('URLTable', () => {
  it('calls onSort when header clicked', () => {
    const onSort = jest.fn()
    render(
      <MemoryRouter>
        <URLTable items={items} selected={[]} onSelect={() => {}} onSort={onSort} sortKey="title" sortDir="asc" />
      </MemoryRouter>
    )
    fireEvent.click(screen.getByText('HTML'))
    expect(onSort).toHaveBeenCalledWith('html_version')
  })

  it('allows selecting items', () => {
    const onSelect = jest.fn()
    render(
      <MemoryRouter>
        <URLTable items={items} selected={[]} onSelect={onSelect} onSort={() => {}} sortKey="title" sortDir="asc" />
      </MemoryRouter>
    )
    fireEvent.click(screen.getByRole('checkbox'))
    expect(onSelect).toHaveBeenCalledWith(1, true)
  })
})
