import clsx from 'clsx'

const LoadingScreen: React.FC = () => {
  return (
    <div
      aria-label="Loading..."
      className={clsx(
        'fixed inset-0 flex items-center justify-center bg-gray-400 bg-opacity-80 z-50',
      )}
      role="status">
      <div
        aria-hidden="true"
        className="h-12 w-12 border-4 border-t-transparent border-white rounded-full animate-spin"
      />
    </div>
  )
}

export default LoadingScreen
