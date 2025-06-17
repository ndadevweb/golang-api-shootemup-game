import { useEffect } from 'react'
import { useForm } from 'react-hook-form'
import { z } from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
import { SigninFormInterface } from '@/features/auth/types/types'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { signin } from '@/features/auth/api/signin'
import { Alert, AlertDescription } from '@/components/ui/alert'
import { AlertCircleIcon } from 'lucide-react'
import useAuthStore from '@/state/useAuthStore'


const signinSchema = z.object({
    login: z.string().min(2, "Identifiant trop court"),
    password: z.string().min(6, "Mot de passe trop court"),
});

// type SigninFormValues = z.infer<typeof signinSchema>

export function SigninForm() {
    const setJwt = useAuthStore((state) => state.setJwt)
    const { register, setFocus, handleSubmit, setError, formState: { errors } } = useForm<SigninFormInterface>({
        resolver: zodResolver(signinSchema),
    })

    const onSubmit = async (data: SigninFormInterface) => {
        try {
            const response = await signin(data);
            setJwt(response.token);
        } catch (error: any) {
            setError('root', { message: error.message });
        }
    }

    useEffect(() => {
        setFocus('login')
    }, [setFocus])

    return (
        <form className="space-y-4 max-w-sm mx-auto mt-8" onSubmit={handleSubmit(onSubmit)}>
            {errors.root && (
                <div className="grid w-full max-w-xl items-start gap-4">

                    <Alert variant="destructive">
                        <AlertCircleIcon />
                        <AlertDescription>
                            <p>{errors.root.message}</p>
                        </AlertDescription>
                    </Alert>
                </div>
            )}
            <div>
                <label htmlFor="login" className="block mb-1">
                    Identifiant
                </label>
                <Input
                    id="login"
                    type="text"
                    placeholder="Entrez votre identifiant"
                    {...register("login")}
                    required
                />
            </div>
            <div>
                <label htmlFor="password" className="block mb-1">
                    Mot de passe
                </label>
                <Input
                    id="password"
                    type="password"
                    placeholder="Votre mot de passe"
                    {...register("password")}
                    required
                />
            </div>
            <Button type="submit" className="w-full">
                Me connecter
            </Button>
        </form>
    )
}