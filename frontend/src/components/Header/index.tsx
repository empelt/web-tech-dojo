import { Button } from "@/components/ui/button";

export const Header = () => {
  return (
    <div className="fixed flex justify-between px-8 w-screen h-16 items-center drop-shadow-2xl shadow-sm">
      <h1 className="font-bold text-2xl">Web技術道場</h1>
      <div className="flex gap-3">
        <Button>ログイン</Button>
      </div>
    </div>
  );
};