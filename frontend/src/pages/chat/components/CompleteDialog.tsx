import { useNavigate } from 'react-router'

import { Button } from '@/components/ui/button'
import {
  Dialog,
  DialogFooter,
  DialogContent,
  DialogTitle,
} from '@/components/ui/dialog'

type Props = {
  onClose: () => void
  open: boolean
}

const CompleteDialog = ({ open, onClose }: Props) => {
  const navigate = useNavigate()

  const handleBackToList = () => {
    navigate('/questions')
    onClose()
  }

  return (
    <Dialog open={open}>
      <DialogContent className="sm:max-w-md" hideCloseButton>
        <DialogTitle>満点です！お疲れ様でした！🎉</DialogTitle>
        <DialogFooter>
          <Button onClick={onClose} variant="outline">
            閉じる
          </Button>
          <Button color="primary" onClick={handleBackToList}>
            問題一覧に戻る
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}

export default CompleteDialog
