import React, { useEffect, useRef } from 'react'

const StarField: React.FC = () => {
  const canvasRef = useRef<HTMLCanvasElement>(null)

  useEffect(() => {
    const canvas = canvasRef.current!
    const ctx = canvas.getContext('2d')!
    let width = window.innerWidth
    let height = window.innerHeight
    canvas.width = width
    canvas.height = height

    const stars = Array.from({ length: 100 }, () => ({
      x: Math.random() * width,
      y: Math.random() * height,
      speed: 0.2 + Math.random() * 0.8,
      radius: 0.5 + Math.random() * 1.5,
    }))

    const draw = () => {
      ctx.clearRect(0, 0, width, height)
      ctx.fillStyle = '#fff'
      ctx.shadowBlur = 2
      ctx.shadowColor = '#fff'
      stars.forEach(star => {
        star.y += star.speed
        if (star.y > height) {
          star.y = 0
          star.x = Math.random() * width
        }
        ctx.beginPath()
        ctx.arc(star.x, star.y, star.radius, 0, Math.PI * 2)
        ctx.fill()
      })
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
        width: '100%',
        height: '100%',
        zIndex: -2,
        pointerEvents: 'none',
      }}
    />
  )
}

export default StarField
