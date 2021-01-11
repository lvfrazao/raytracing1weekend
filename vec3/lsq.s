// These are all alternate implementations of LengthSquared in ASM
// These are all slower since they can't be inlined so they spec a ton
// of time putting and taking stuff on the stack

// This is basically the same implementation as the pure Go function 
TEXT ·lensq(SB), 4, $0 
    // Grab function parameters off the stack
    MOVSD  x+0(FP),X0
    MOVSD  y+8(FP),X1
    MOVSD  z+16(FP),X2

    MULSD  X1,X1
    MULSD  X0,X0
    MULSD  X2,X2
    ADDSD  X1,X0
    ADDSD  X2,X0

    // Place return val in correct location on stack
    MOVSD  X0,ret+24(FP)
    RET

TEXT ·lensq2(SB), $0
    MOVSD  x+0(FP),X0
    MOVSD  y+8(FP),X1
    MOVSD  z+16(FP),X2

    // x2 = x2*x2
    VMULSD X2,X2,X2
    // x1 = (x1 * x1) + x2
    VFMADD132SD X1,X2,X1
    // x0 = (x0 * x0) + x1
    VFMADD132SD X0,X1,X0
    
    MOVSD  X0,ret+24(FP)
    RET

TEXT ·lensq3(SB), 4, $0 
    MOVSD  x+0(FP),X0
    MOVSD  y+8(FP),X1
    MOVSD  z+16(FP),X2

    VMOVQ  X2,X2
    VUNPCKLPD X1,X0,X1
    VINSERTF128 $0x1,X2,Y1,Y1
    VMULPD Y1,Y1,Y1
    VHADDPD X1,X1,X0
    VEXTRACTF128 $0x1,Y1,X1
    VADDSD X1,X0,X0
    VZEROUPPER 

    MOVSD  X0,ret+24(FP)
    RET

TEXT ·lensq4(SB), 4, $0 
    MOVSD  x+0(FP),X0
    MOVSD  y+8(FP),X1
    MOVSD  z+16(FP),X2

    VMOVQ  X2,X2
    VUNPCKLPD X1,X0,X0
    VINSERTF128 $0x1,X2,Y0,Y0
    VMULPD Y0,Y0,Y0
    VMOVAPS X0,X1
    VEXTRACTF128 $0x1,Y0,X0
    VADDPD X0,X1,X0
    VUNPCKHPD X0,X0,X1
    VADDSD X1,X0,X0
    VZEROUPPER 

    MOVSD  X0,ret+24(FP)
    RET

TEXT ·lensq5(SB), 4, $0 
    MOVSD  x+0(FP),X0
    MOVSD  y+8(FP),X1
    MOVSD  z+16(FP),X2

    VMOVQ  X2,X2
    VUNPCKLPD X1,X0,X0
    VMULPD X2,X2,X2
    VFMADD132PD X0,X2,X0
    VUNPCKHPD X0,X0,X1
    VADDSD X1,X0,X0

    MOVSD  X0,ret+24(FP)
    RET
