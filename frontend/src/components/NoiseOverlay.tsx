import React, { useEffect, useRef } from 'react'

const NoiseOverlay: React.FC = () => {
  const canvasRef = useRef<HTMLCanvasElement>(null)

  useEffect(() => {
    const canvas = canvasRef.current!
    const ctx = canvas.getContext('2d')!
    let width = window.innerWidth
    let height = window.innerHeight
    canvas.width = width
    canvas.height = height

    const draw = () => {
      const id = ctx.createImageData(width, height)
      for (let i = 0; i < id.data.length; i += 4) {
        const shade = Math.random() * 255
        id.data[i] = 100 + shade * 0.6
        id.data[i + 1] = shade * 0.2
        id.data[i + 2] = 150 + shade * 0.6
        id.data[i + 3] = 20
      }
      ctx.putImageData(id, 0, 0)
    }

    let frame: number
    const step = () => {
      draw()
      frame = requestAnimationFrame(step)
    }
    step()

    const handleResize = () => {
      width = window.innerWidth
      height = window.innerHeight
      canvas.width = width
      canvas.height = height
    }
    window.addEventListener('resize', handleResize)

    return () => {
      cancelAnimationFrame(frame)
      window.removeEventListener('resize', handleResize)
    }
  }, [])

  return (
    <canvas
      ref={canvasRef}
      style={{
        position: 'fixed',
        top: 0,
        left: 0,
        zIndex: -1,
        pointerEvents: 'none',
        mixBlendMode: 'screen',
        filter: 'blur(1px)',
      }}
    />
  )
}

export default NoiseOverlay
