<?php

namespace App\Http\Controllers\Api;

use Auth;
use App\User;
use App\Http\Requests\Api\LoginUser;
use Illuminate\Support\Facades\Redis;
use App\Http\Requests\Api\RegisterUser;
use App\Http\Controllers\Api\UserController;
use App\RealWorld\Transformers\UserTransformer;


class AuthController extends ApiController
{
    /**
     * AuthController constructor.
     *
     * @param UserTransformer $transformer
     */
    public function __construct(UserTransformer $transformer)
    {
        $this->transformer = $transformer;
    }

    /**
     * Login user and return the user if successful.
     *
     * @param LoginUser $request
     * @return \Illuminate\Http\JsonResponse
     */
    public function login(LoginUser $request)
    {
        $credentials = $request->only('user.email', 'user.password');
        $credentials = $credentials['user'];

        $redis = json_decode(Redis::get($credentials['email']), true);

        $userAdmin = UserController::show($credentials['email']);

        $code = $redis['code'];

        // Clave 5º
        $clv5 = substr($userAdmin['bearer'], 0, 5);
        $code = substr($code, 0, -5);

        //Clave 4º
        $clv4 = substr($userAdmin['username'], -2);
        $code = substr($code, 0, -2);

        //Clave 3º
        $clv3 = substr((explode('@', $userAdmin['email'])[0]), -3);
        $code = substr($code, 0, -3);

        //Clave 2º
        $clv2 = substr($userAdmin['username'], 0, 3);
        $code = substr($code, 0, -3);

        $code_lar = $code . $clv2 . $clv3 . $clv4 . $clv5;

        return (($userAdmin['bearer'] === $redis['bearer']) && ($code_lar === $redis['code']) && (Auth::once($credentials))) ?
        
            $this->respondWithTransformer(auth()->user()) :

            $this->respondFailedLogin();
    }


    /**
     * Register a new user and return the user if successful.
     *
     * @param RegisterUser $request
     * @return \Illuminate\Http\JsonResponse
     */
    public function register(RegisterUser $request)
    {
        $user = User::create([
            'username' => $request->input('user.username'),
            'email' => $request->input('user.email'),
            'password' => bcrypt($request->input('user.password')),
        ]);

        return $this->respondWithTransformer($user);
    }
}
