(module
  (type (;0;) (func (param i32 i32 i32 i32) (result i32)))
  (type (;1;) (func (param i32 i32) (result i32)))
  (type (;2;) (func (param i32 i64 i32) (result i32)))
  (type (;3;) (func (param i32 i32)))
  (type (;4;) (func))
  (type (;5;) (func (param i32) (result i32)))
  (type (;6;) (func (param i32)))
  (type (;7;) (func (param i32 i32 i32) (result i32)))
  (import "wasi_snapshot_preview1" "fd_write" (func $runtime.fd_write (type 0)))
  (import "wasi_snapshot_preview1" "args_sizes_get" (func $runtime.args_sizes_get (type 1)))
  (import "wasi_snapshot_preview1" "args_get" (func $runtime.args_get (type 1)))
  (import "wasi_snapshot_preview1" "clock_time_get" (func $runtime.clock_time_get (type 2)))
  (import "temphia" "log" (func $github.com/temphia/temphia/core/backend/server/engine/executors/wasm1/sdk._log (type 3)))
  (func $__wasm_call_ctors (type 4))
  (func $runtime.alloc (type 5) (param i32) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32)
    block  ;; label = @1
      local.get 0
      br_if 0 (;@1;)
      i32.const 65812
      return
    end
    local.get 0
    i32.const 15
    i32.add
    i32.const 4
    i32.shr_u
    local.set 1
    i32.const 66064
    i32.const -4
    i32.and
    local.set 2
    i32.const 0
    i32.load offset=65804
    local.tee 3
    local.set 4
    i32.const 0
    local.set 5
    i32.const 0
    local.set 6
    loop (result i32)  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 4
          local.get 3
          i32.ne
          br_if 0 (;@3;)
          i32.const 1
          local.set 7
          block  ;; label = @4
            block  ;; label = @5
              local.get 6
              i32.const 255
              i32.and
              br_table 3 (;@2;) 0 (;@5;) 1 (;@4;)
            end
            i32.const 65816
            local.set 7
            block  ;; label = @5
              loop  ;; label = @6
                local.get 7
                i32.load
                local.tee 7
                i32.eqz
                br_if 1 (;@5;)
                local.get 7
                i32.const 8
                i32.add
                local.tee 6
                local.get 6
                local.get 7
                i32.load offset=4
                i32.const 2
                i32.shl
                i32.add
                call $runtime.markRoots
                br 0 (;@6;)
              end
            end
            i32.const 65536
            local.get 2
            call $runtime.markRoots
            loop  ;; label = @5
              block  ;; label = @6
                i32.const 0
                i32.load8_u offset=65813
                br_if 0 (;@6;)
                i32.const 0
                local.set 6
                i32.const 0
                local.set 7
                loop  ;; label = @7
                  block  ;; label = @8
                    local.get 7
                    i32.const 0
                    i32.load offset=65808
                    i32.lt_u
                    br_if 0 (;@8;)
                    i32.const 2
                    local.set 7
                    br 6 (;@2;)
                  end
                  block  ;; label = @8
                    block  ;; label = @9
                      block  ;; label = @10
                        block  ;; label = @11
                          local.get 7
                          call $_runtime.gcBlock_.state
                          i32.const 255
                          i32.and
                          i32.const -1
                          i32.add
                          br_table 1 (;@10;) 0 (;@11;) 2 (;@9;) 3 (;@8;)
                        end
                        local.get 6
                        i32.const 1
                        i32.and
                        local.set 3
                        i32.const 0
                        local.set 6
                        local.get 3
                        i32.eqz
                        br_if 2 (;@8;)
                      end
                      local.get 7
                      call $_runtime.gcBlock_.markFree
                      i32.const 1
                      local.set 6
                      br 1 (;@8;)
                    end
                    i32.const 0
                    local.set 6
                    i32.const 0
                    i32.load offset=65800
                    local.get 7
                    i32.const 2
                    i32.shr_u
                    i32.add
                    local.tee 3
                    local.get 3
                    i32.load8_u
                    i32.const 2
                    local.get 7
                    i32.const 1
                    i32.shl
                    i32.const 6
                    i32.and
                    i32.shl
                    i32.const -1
                    i32.xor
                    i32.and
                    i32.store8
                  end
                  local.get 7
                  i32.const 1
                  i32.add
                  local.set 7
                  br 0 (;@7;)
                end
              end
              i32.const 0
              local.set 7
              i32.const 0
              i32.const 0
              i32.store8 offset=65813
              loop  ;; label = @6
                local.get 7
                i32.const 0
                i32.load offset=65808
                i32.ge_u
                br_if 1 (;@5;)
                block  ;; label = @7
                  local.get 7
                  call $_runtime.gcBlock_.state
                  i32.const 255
                  i32.and
                  i32.const 3
                  i32.ne
                  br_if 0 (;@7;)
                  local.get 7
                  call $runtime.startMark
                end
                local.get 7
                i32.const 1
                i32.add
                local.set 7
                br 0 (;@6;)
              end
            end
          end
          block  ;; label = @4
            memory.size
            memory.grow
            i32.const -1
            i32.eq
            br_if 0 (;@4;)
            memory.size
            local.set 7
            i32.const 0
            i32.load offset=65672
            local.set 3
            i32.const 0
            local.get 7
            i32.const 16
            i32.shl
            i32.store offset=65672
            i32.const 0
            i32.load offset=65800
            local.set 7
            call $runtime.calculateHeapAddresses
            i32.const 0
            i32.load offset=65800
            local.get 7
            local.get 3
            local.get 7
            i32.sub
            call $memcpy
            drop
            br 1 (;@3;)
          end
          i32.const 65536
          i32.const 13
          call $runtime.runtimePanic
          unreachable
        end
        local.get 6
        local.set 7
      end
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            local.get 4
            i32.const 0
            i32.load offset=65808
            i32.ne
            br_if 0 (;@4;)
            i32.const 0
            local.set 4
            br 1 (;@3;)
          end
          block  ;; label = @4
            local.get 4
            call $_runtime.gcBlock_.state
            i32.const 255
            i32.and
            i32.eqz
            br_if 0 (;@4;)
            local.get 4
            i32.const 1
            i32.add
            local.set 4
            br 1 (;@3;)
          end
          local.get 4
          i32.const 1
          i32.add
          local.set 6
          block  ;; label = @4
            local.get 5
            i32.const 1
            i32.add
            local.tee 5
            local.get 1
            i32.eq
            br_if 0 (;@4;)
            local.get 6
            local.set 4
            br 2 (;@2;)
          end
          i32.const 0
          local.get 6
          i32.store offset=65804
          local.get 6
          local.get 1
          i32.sub
          local.tee 6
          i32.const 1
          call $_runtime.gcBlock_.setState
          local.get 4
          local.get 1
          i32.sub
          i32.const 2
          i32.add
          local.set 7
          block  ;; label = @4
            loop  ;; label = @5
              local.get 7
              i32.const 0
              i32.load offset=65804
              i32.eq
              br_if 1 (;@4;)
              local.get 7
              i32.const 2
              call $_runtime.gcBlock_.setState
              local.get 7
              i32.const 1
              i32.add
              local.set 7
              br 0 (;@5;)
            end
          end
          i32.const 0
          i32.load offset=65632
          local.get 6
          i32.const 4
          i32.shl
          i32.add
          i32.const 0
          local.get 0
          call $memset
          return
        end
        i32.const 0
        local.set 5
      end
      i32.const 0
      i32.load offset=65804
      local.set 3
      local.get 7
      local.set 6
      br 0 (;@1;)
    end)
  (func $runtime.markRoots (type 3) (param i32 i32)
    (local i32)
    block  ;; label = @1
      loop  ;; label = @2
        local.get 0
        local.get 1
        i32.ge_u
        br_if 1 (;@1;)
        block  ;; label = @3
          local.get 0
          i32.load
          local.tee 2
          call $runtime.looksLikePointer
          i32.const 1
          i32.and
          i32.eqz
          br_if 0 (;@3;)
          local.get 2
          i32.const 0
          i32.load offset=65632
          i32.sub
          i32.const 4
          i32.shr_u
          local.tee 2
          call $_runtime.gcBlock_.state
          i32.const 255
          i32.and
          i32.eqz
          br_if 0 (;@3;)
          local.get 2
          call $_runtime.gcBlock_.findHead
          local.tee 2
          call $_runtime.gcBlock_.state
          i32.const 255
          i32.and
          i32.const 3
          i32.eq
          br_if 0 (;@3;)
          local.get 2
          call $runtime.startMark
        end
        local.get 0
        i32.const 4
        i32.add
        local.set 0
        br 0 (;@2;)
      end
    end)
  (func $_runtime.gcBlock_.state (type 5) (param i32) (result i32)
    i32.const 0
    i32.load offset=65800
    local.get 0
    i32.const 2
    i32.shr_u
    i32.add
    i32.load8_u
    local.get 0
    i32.const 1
    i32.shl
    i32.const 6
    i32.and
    i32.shr_u
    i32.const 3
    i32.and)
  (func $_runtime.gcBlock_.markFree (type 6) (param i32)
    (local i32)
    i32.const 0
    i32.load offset=65800
    local.get 0
    i32.const 2
    i32.shr_u
    i32.add
    local.tee 1
    local.get 1
    i32.load8_u
    i32.const 3
    local.get 0
    i32.const 1
    i32.shl
    i32.const 6
    i32.and
    i32.shl
    i32.const -1
    i32.xor
    i32.and
    i32.store8)
  (func $runtime.startMark (type 6) (param i32)
    (local i32 i32 i32 i32)
    global.get $__stack_pointer
    i32.const 64
    i32.sub
    local.tee 1
    global.set $__stack_pointer
    local.get 1
    i32.const 0
    i32.const 64
    call $memset
    local.tee 2
    local.get 0
    i32.store
    local.get 0
    i32.const 3
    call $_runtime.gcBlock_.setState
    i32.const 1
    local.set 3
    block  ;; label = @1
      loop  ;; label = @2
        local.get 3
        i32.const 1
        i32.lt_s
        br_if 1 (;@1;)
        block  ;; label = @3
          local.get 3
          i32.const -1
          i32.add
          local.tee 3
          i32.const 15
          i32.gt_u
          br_if 0 (;@3;)
          i32.const 0
          i32.load offset=65632
          local.get 2
          local.get 3
          i32.const 2
          i32.shl
          i32.add
          i32.load
          local.tee 1
          i32.const 4
          i32.shl
          i32.add
          local.set 0
          local.get 1
          call $_runtime.gcBlock_.findNext
          local.set 1
          i32.const 0
          i32.load offset=65632
          local.get 1
          i32.const 4
          i32.shl
          i32.add
          local.set 4
          loop  ;; label = @4
            local.get 4
            local.get 0
            i32.eq
            br_if 2 (;@2;)
            block  ;; label = @5
              local.get 0
              i32.load
              local.tee 1
              call $runtime.looksLikePointer
              i32.const 1
              i32.and
              i32.eqz
              br_if 0 (;@5;)
              local.get 1
              i32.const 0
              i32.load offset=65632
              i32.sub
              i32.const 4
              i32.shr_u
              local.tee 1
              call $_runtime.gcBlock_.state
              i32.const 255
              i32.and
              i32.eqz
              br_if 0 (;@5;)
              local.get 1
              call $_runtime.gcBlock_.findHead
              local.tee 1
              call $_runtime.gcBlock_.state
              i32.const 255
              i32.and
              i32.const 3
              i32.eq
              br_if 0 (;@5;)
              local.get 1
              i32.const 3
              call $_runtime.gcBlock_.setState
              block  ;; label = @6
                local.get 3
                i32.const 16
                i32.ne
                br_if 0 (;@6;)
                i32.const 0
                i32.const 1
                i32.store8 offset=65813
                i32.const 16
                local.set 3
                br 1 (;@5;)
              end
              local.get 3
              i32.const 15
              i32.gt_u
              br_if 2 (;@3;)
              local.get 2
              local.get 3
              i32.const 2
              i32.shl
              i32.add
              local.get 1
              i32.store
              local.get 3
              i32.const 1
              i32.add
              local.set 3
            end
            local.get 0
            i32.const 4
            i32.add
            local.set 0
            br 0 (;@4;)
          end
        end
      end
      call $runtime.lookupPanic
      unreachable
    end
    local.get 2
    i32.const 64
    i32.add
    global.set $__stack_pointer)
  (func $runtime.calculateHeapAddresses (type 4)
    (local i32 i32)
    i32.const 0
    i32.const 0
    i32.load offset=65632
    i32.const 15
    i32.add
    i32.const -16
    i32.and
    local.tee 0
    i32.store offset=65632
    i32.const 0
    i32.const 0
    i32.load offset=65672
    local.tee 1
    local.get 1
    local.get 0
    i32.sub
    i32.const 64
    i32.add
    i32.const 65
    i32.div_u
    i32.sub
    local.tee 1
    i32.store offset=65800
    i32.const 0
    local.get 1
    local.get 0
    i32.sub
    i32.const 4
    i32.shr_u
    i32.store offset=65808)
  (func $runtime.runtimePanic (type 3) (param i32 i32)
    i32.const 65549
    i32.const 22
    call $runtime.printstring
    local.get 0
    local.get 1
    call $runtime.printstring
    call $runtime.printnl
    unreachable
    unreachable)
  (func $_runtime.gcBlock_.setState (type 3) (param i32 i32)
    (local i32)
    i32.const 0
    i32.load offset=65800
    local.get 0
    i32.const 2
    i32.shr_u
    i32.add
    local.tee 2
    local.get 2
    i32.load8_u
    local.get 1
    local.get 0
    i32.const 1
    i32.shl
    i32.const 6
    i32.and
    i32.shl
    i32.or
    i32.store8)
  (func $runtime.printstring (type 3) (param i32 i32)
    local.get 1
    i32.const 0
    local.get 1
    i32.const 0
    i32.gt_s
    select
    local.set 1
    block  ;; label = @1
      loop  ;; label = @2
        local.get 1
        i32.eqz
        br_if 1 (;@1;)
        local.get 0
        i32.load8_u
        call $runtime.putchar
        local.get 0
        i32.const 1
        i32.add
        local.set 0
        local.get 1
        i32.const -1
        i32.add
        local.set 1
        br 0 (;@2;)
      end
    end)
  (func $runtime.printnl (type 4)
    i32.const 10
    call $runtime.putchar)
  (func $runtime.putchar (type 6) (param i32)
    (local i32 i32)
    block  ;; label = @1
      i32.const 0
      i32.load offset=65676
      local.tee 1
      i32.const 119
      i32.le_u
      br_if 0 (;@1;)
      call $runtime.lookupPanic
      unreachable
    end
    i32.const 0
    local.get 1
    i32.const 1
    i32.add
    local.tee 2
    i32.store offset=65676
    local.get 1
    i32.const 65680
    i32.add
    local.get 0
    i32.store8
    block  ;; label = @1
      block  ;; label = @2
        local.get 0
        i32.const 255
        i32.and
        i32.const 10
        i32.eq
        br_if 0 (;@2;)
        local.get 1
        i32.const 119
        i32.ne
        br_if 1 (;@1;)
      end
      i32.const 0
      local.get 2
      i32.store offset=65640
      i32.const 1
      i32.const 65636
      i32.const 1
      i32.const 65820
      call $runtime.fd_write
      drop
      i32.const 0
      i32.const 0
      i32.store offset=65676
    end)
  (func $runtime.lookupPanic (type 4)
    i32.const 65594
    i32.const 18
    call $runtime.runtimePanic
    unreachable)
  (func $_runtime.gcBlock_.findNext (type 5) (param i32) (result i32)
    (local i32)
    block  ;; label = @1
      block  ;; label = @2
        local.get 0
        call $_runtime.gcBlock_.state
        i32.const 255
        i32.and
        i32.const 1
        i32.eq
        br_if 0 (;@2;)
        local.get 0
        call $_runtime.gcBlock_.state
        i32.const 255
        i32.and
        i32.const 3
        i32.ne
        br_if 1 (;@1;)
      end
      local.get 0
      i32.const 1
      i32.add
      local.set 0
    end
    local.get 0
    i32.const 4
    i32.shl
    local.set 1
    block  ;; label = @1
      loop  ;; label = @2
        local.get 1
        i32.const 0
        i32.load offset=65632
        i32.add
        i32.const 0
        i32.load offset=65800
        i32.ge_u
        br_if 1 (;@1;)
        local.get 0
        call $_runtime.gcBlock_.state
        i32.const 255
        i32.and
        i32.const 2
        i32.ne
        br_if 1 (;@1;)
        local.get 1
        i32.const 16
        i32.add
        local.set 1
        local.get 0
        i32.const 1
        i32.add
        local.set 0
        br 0 (;@2;)
      end
    end
    local.get 0)
  (func $runtime.looksLikePointer (type 5) (param i32) (result i32)
    (local i32)
    i32.const 0
    local.set 1
    block  ;; label = @1
      i32.const 0
      i32.load offset=65632
      local.get 0
      i32.gt_u
      br_if 0 (;@1;)
      i32.const 0
      i32.load offset=65800
      local.get 0
      i32.gt_u
      local.set 1
    end
    local.get 1)
  (func $_runtime.gcBlock_.findHead (type 5) (param i32) (result i32)
    (local i32 i32)
    loop  ;; label = @1
      local.get 0
      call $_runtime.gcBlock_.state
      local.set 1
      local.get 0
      i32.const -1
      i32.add
      local.tee 2
      local.set 0
      local.get 1
      i32.const 255
      i32.and
      i32.const 2
      i32.eq
      br_if 0 (;@1;)
    end
    local.get 2
    i32.const 1
    i32.add)
  (func $runtime.nilPanic (type 4)
    i32.const 65571
    i32.const 23
    call $runtime.runtimePanic
    unreachable)
  (func $runtime.hash32 (type 7) (param i32 i32 i32) (result i32)
    (local i32)
    i32.const -2128831035
    local.set 3
    block  ;; label = @1
      loop  ;; label = @2
        local.get 1
        i32.eqz
        br_if 1 (;@1;)
        local.get 1
        i32.const -1
        i32.add
        local.set 1
        local.get 3
        local.get 0
        i32.load8_u
        i32.xor
        i32.const 16777619
        i32.mul
        local.set 3
        local.get 0
        i32.const 1
        i32.add
        local.set 0
        br 0 (;@2;)
      end
    end
    local.get 3)
  (func $malloc (type 5) (param i32) (result i32)
    (local i32 i32)
    global.get $__stack_pointer
    i32.const 16
    i32.sub
    local.tee 1
    global.set $__stack_pointer
    local.get 1
    i64.const 1
    i64.store offset=4 align=4
    i32.const 0
    i32.load offset=65816
    local.set 2
    i32.const 0
    local.get 1
    i32.store offset=65816
    local.get 1
    local.get 2
    i32.store
    local.get 0
    call $runtime.alloc
    local.set 0
    i32.const 0
    local.get 2
    i32.store offset=65816
    local.get 1
    i32.const 8
    i32.add
    local.get 0
    i32.store
    local.get 1
    i32.const 16
    i32.add
    global.set $__stack_pointer
    local.get 0)
  (func $free (type 6) (param i32))
  (func $calloc (type 1) (param i32 i32) (result i32)
    (local i32 i32)
    global.get $__stack_pointer
    i32.const 16
    i32.sub
    local.tee 2
    global.set $__stack_pointer
    local.get 2
    i64.const 1
    i64.store offset=4 align=4
    i32.const 0
    i32.load offset=65816
    local.set 3
    i32.const 0
    local.get 2
    i32.store offset=65816
    local.get 2
    local.get 3
    i32.store
    local.get 1
    local.get 0
    i32.mul
    call $runtime.alloc
    local.set 1
    i32.const 0
    local.get 3
    i32.store offset=65816
    local.get 2
    i32.const 8
    i32.add
    local.get 1
    i32.store
    local.get 2
    i32.const 16
    i32.add
    global.set $__stack_pointer
    local.get 1)
  (func $realloc (type 1) (param i32 i32) (result i32)
    (local i32 i32 i32)
    global.get $__stack_pointer
    i32.const 32
    i32.sub
    local.tee 2
    global.set $__stack_pointer
    local.get 2
    i64.const 0
    i64.store offset=20 align=4
    local.get 2
    i64.const 3
    i64.store offset=12 align=4
    i32.const 0
    i32.load offset=65816
    local.set 3
    i32.const 0
    local.get 2
    i32.const 8
    i32.add
    i32.store offset=65816
    local.get 2
    local.get 3
    i32.store offset=8
    block  ;; label = @1
      block  ;; label = @2
        local.get 0
        br_if 0 (;@2;)
        local.get 2
        i32.const 16
        i32.add
        local.get 1
        call $runtime.alloc
        local.tee 1
        i32.store
        br 1 (;@1;)
      end
      block  ;; label = @2
        local.get 0
        i32.const 0
        i32.load offset=65632
        i32.sub
        i32.const 4
        i32.shr_u
        call $_runtime.gcBlock_.findNext
        i32.const 4
        i32.shl
        local.get 0
        i32.sub
        i32.const 0
        i32.load offset=65632
        i32.add
        local.tee 4
        local.get 1
        i32.lt_u
        br_if 0 (;@2;)
        local.get 0
        local.set 1
        br 1 (;@1;)
      end
      local.get 2
      i32.const 20
      i32.add
      local.get 1
      call $runtime.alloc
      local.tee 1
      i32.store
      local.get 1
      local.get 0
      local.get 4
      call $memcpy
      drop
    end
    i32.const 0
    local.get 3
    i32.store offset=65816
    local.get 2
    i32.const 24
    i32.add
    local.get 1
    i32.store
    local.get 2
    i32.const 32
    i32.add
    global.set $__stack_pointer
    local.get 1)
  (func $runtime.slicePanic (type 4)
    i32.const 65612
    i32.const 18
    call $runtime.runtimePanic
    unreachable)
  (func $runtime.memequal (type 0) (param i32 i32 i32 i32) (result i32)
    (local i32 i32)
    i32.const 0
    local.set 4
    block  ;; label = @1
      loop  ;; label = @2
        local.get 2
        local.get 4
        local.tee 5
        i32.eq
        br_if 1 (;@1;)
        local.get 5
        i32.const 1
        i32.add
        local.set 4
        local.get 0
        local.get 5
        i32.add
        i32.load8_u
        local.get 1
        local.get 5
        i32.add
        i32.load8_u
        i32.eq
        br_if 0 (;@2;)
      end
    end
    local.get 5
    local.get 2
    i32.ge_u)
  (func $_start (type 4)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32)
    global.get $__stack_pointer
    i32.const 192
    i32.sub
    local.tee 0
    global.set $__stack_pointer
    i32.const 0
    local.set 1
    i32.const 0
    i32.const 66064
    i32.store offset=65632
    local.get 0
    i32.const 152
    i32.add
    local.tee 2
    i64.const 0
    i64.store
    local.get 0
    i32.const 112
    i32.add
    local.tee 3
    i64.const 0
    i64.store
    local.get 0
    i32.const 88
    i32.add
    i32.const 16
    i32.add
    local.tee 4
    i64.const 0
    i64.store
    local.get 0
    i32.const 96
    i32.add
    local.tee 5
    i64.const 0
    i64.store
    local.get 0
    i64.const 98784247808
    i64.store offset=88
    local.get 0
    i32.const 0
    i32.store offset=184
    local.get 0
    i64.const 0
    i64.store offset=176
    local.get 0
    i64.const 0
    i64.store offset=168
    local.get 0
    i64.const 0
    i64.store offset=160
    local.get 0
    i64.const 0
    i64.store offset=144
    local.get 0
    i64.const 0
    i64.store offset=136
    local.get 0
    i64.const 0
    i64.store offset=128
    local.get 0
    i64.const 0
    i64.store offset=120
    i32.const 0
    i32.load offset=65816
    local.set 6
    i32.const 0
    local.get 0
    i32.const 88
    i32.add
    i32.store offset=65816
    local.get 0
    local.get 6
    i32.store offset=88
    i32.const 0
    memory.size
    i32.const 16
    i32.shl
    i32.store offset=65672
    call $runtime.calculateHeapAddresses
    local.get 5
    i32.const 0
    i32.load offset=65800
    local.tee 7
    i32.store
    local.get 0
    i32.const 100
    i32.add
    local.get 7
    i32.store
    local.get 7
    i32.const 0
    i32.const 0
    i32.load offset=65672
    local.get 7
    i32.sub
    call $memset
    drop
    i32.const 0
    memory.size
    i32.const 16
    i32.shl
    i32.store offset=65672
    call $__wasm_call_ctors
    local.get 4
    i32.const 0
    i32.load offset=65824
    local.tee 7
    i32.store
    local.get 0
    i32.const 108
    i32.add
    local.get 0
    i32.const 32
    i32.add
    i32.store
    local.get 3
    local.get 0
    i32.const 8
    i32.add
    i32.store
    local.get 0
    i32.const 156
    i32.add
    local.get 0
    i32.const 64
    i32.add
    i32.store
    local.get 2
    local.get 0
    i32.const 64
    i32.add
    i32.store
    local.get 0
    i32.const 140
    i32.add
    local.get 0
    i32.const 64
    i32.add
    i32.store
    block  ;; label = @1
      block  ;; label = @2
        local.get 7
        br_if 0 (;@2;)
        local.get 0
        i32.const 0
        i32.store offset=8
        local.get 0
        i32.const 0
        i32.store offset=32
        local.get 0
        i32.const 32
        i32.add
        local.get 0
        i32.const 8
        i32.add
        call $runtime.args_sizes_get
        drop
        block  ;; label = @3
          local.get 0
          i32.load offset=32
          local.tee 8
          br_if 0 (;@3;)
          i32.const 0
          local.set 7
          i32.const 0
          local.set 2
          br 2 (;@1;)
        end
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              local.get 8
              i32.const 1073741823
              i32.gt_u
              br_if 0 (;@5;)
              local.get 0
              i32.const 120
              i32.add
              local.get 8
              i32.const 2
              i32.shl
              call $runtime.alloc
              local.tee 2
              i32.store
              local.get 0
              i32.load offset=8
              local.tee 7
              i32.const -1
              i32.le_s
              br_if 0 (;@5;)
              local.get 0
              i32.const 124
              i32.add
              local.get 7
              call $runtime.alloc
              local.tee 1
              i32.store
              local.get 0
              i32.const 128
              i32.add
              local.get 1
              i32.store
              local.get 7
              i32.eqz
              br_if 2 (;@3;)
              local.get 2
              local.get 1
              call $runtime.args_get
              drop
              local.get 8
              i32.const 536870912
              i32.ge_u
              br_if 0 (;@5;)
              local.get 8
              i32.const 3
              i32.shl
              call $runtime.alloc
              local.set 7
              i32.const 0
              local.get 8
              i32.store offset=65832
              i32.const 0
              local.get 8
              i32.store offset=65828
              i32.const 0
              local.get 7
              i32.store offset=65824
              local.get 0
              i32.const 132
              i32.add
              local.get 7
              i32.store
              local.get 0
              i32.const 160
              i32.add
              local.set 9
              local.get 0
              i32.const 144
              i32.add
              local.set 10
              local.get 0
              i32.const 136
              i32.add
              local.set 11
              local.get 0
              i32.const 148
              i32.add
              local.set 12
              i32.const 4
              local.set 3
              i32.const 0
              local.set 1
              br 1 (;@4;)
            end
            call $runtime.slicePanic
            unreachable
          end
          loop  ;; label = @4
            local.get 8
            local.get 1
            i32.eq
            br_if 2 (;@2;)
            local.get 9
            local.get 2
            i32.load
            local.tee 7
            i32.store
            local.get 10
            local.get 7
            i32.store
            local.get 11
            local.get 7
            i32.store
            local.get 7
            call $strlen
            local.set 4
            local.get 12
            i32.const 0
            i32.load offset=65824
            local.tee 5
            i32.store
            local.get 0
            i64.const 0
            i64.store offset=64
            local.get 0
            local.get 7
            i32.store offset=64
            local.get 0
            local.get 4
            i32.store offset=68
            local.get 1
            i32.const 0
            i32.load offset=65828
            i32.ge_u
            br_if 1 (;@3;)
            local.get 5
            local.get 3
            i32.add
            local.tee 5
            local.get 4
            i32.store
            local.get 5
            i32.const -4
            i32.add
            local.get 7
            i32.store
            local.get 2
            i32.const 4
            i32.add
            local.set 2
            local.get 3
            i32.const 8
            i32.add
            local.set 3
            local.get 1
            i32.const 1
            i32.add
            local.set 1
            br 0 (;@4;)
          end
        end
        call $runtime.lookupPanic
        unreachable
      end
      local.get 0
      i32.const 116
      i32.add
      i32.const 0
      i32.load offset=65824
      local.tee 7
      i32.store
      i32.const 0
      i32.load offset=65832
      local.set 2
      i32.const 0
      i32.load offset=65828
      local.set 1
    end
    i32.const 0
    local.get 7
    i32.store offset=65904
    local.get 0
    i32.const 164
    i32.add
    local.get 7
    i32.store
    i32.const 0
    local.get 1
    i32.store offset=65908
    i32.const 0
    local.get 2
    i32.store offset=65912
    local.get 0
    i32.const 168
    i32.add
    local.get 0
    i32.const 8
    i32.add
    i32.store
    local.get 0
    i32.const 8
    i32.add
    i32.const 16
    i32.add
    i32.const 0
    i32.store
    local.get 0
    i32.const 8
    i32.add
    i32.const 8
    i32.add
    i64.const 0
    i64.store
    local.get 0
    i64.const 0
    i64.store offset=8
    local.get 0
    i32.const 172
    i32.add
    local.get 0
    i32.const 56
    i32.add
    i32.store
    local.get 0
    i64.const 0
    i64.store offset=56
    i32.const 0
    i64.const 1000
    local.get 0
    i32.const 56
    i32.add
    call $runtime.clock_time_get
    drop
    local.get 0
    i32.const 176
    i32.add
    local.get 0
    i32.const 64
    i32.add
    local.get 0
    i32.const 32
    i32.add
    local.get 0
    i64.load offset=56
    i64.const 1000000000
    i64.div_s
    i64.const 2682288000
    i64.add
    i64.const 8589934592
    i64.lt_u
    select
    local.tee 7
    i32.store
    local.get 7
    i64.const 0
    i64.store
    local.get 7
    i32.const 8
    i32.add
    i64.const 0
    i64.store
    local.get 7
    i32.const 16
    i32.add
    i32.const 65840
    i32.store
    i32.const 0
    local.get 6
    i32.store offset=65816
    local.get 0
    i32.const 180
    i32.add
    local.get 0
    i32.const 64
    i32.add
    i32.store
    local.get 0
    i32.const 184
    i32.add
    local.get 0
    i32.const 64
    i32.add
    i32.store
    local.get 0
    i32.const 192
    i32.add
    global.set $__stack_pointer)
  (func $action_main (type 3) (param i32 i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32)
    global.get $__stack_pointer
    i32.const 160
    i32.sub
    local.tee 2
    global.set $__stack_pointer
    local.get 2
    i32.const 24
    i32.add
    i32.const 8
    i32.add
    local.tee 3
    i64.const 0
    i64.store
    local.get 2
    i64.const 0
    i64.store offset=152
    local.get 2
    i64.const 0
    i64.store offset=144
    local.get 2
    i64.const 0
    i64.store offset=136
    local.get 2
    i64.const 0
    i64.store offset=128
    local.get 2
    i64.const 0
    i64.store offset=120
    local.get 2
    i64.const 0
    i64.store offset=112
    local.get 2
    i64.const 0
    i64.store offset=104
    local.get 2
    i64.const 0
    i64.store offset=96
    local.get 2
    i64.const 0
    i64.store offset=88
    local.get 2
    i64.const 0
    i64.store offset=80
    local.get 2
    i64.const 0
    i64.store offset=72
    local.get 2
    i64.const 0
    i64.store offset=64
    local.get 2
    i64.const 0
    i64.store offset=56
    local.get 2
    i64.const 0
    i64.store offset=48
    local.get 2
    i64.const 0
    i64.store offset=40
    local.get 2
    i64.const 137438953472
    i64.store offset=24
    local.get 2
    i32.const 0
    i32.load offset=65816
    local.tee 4
    i32.store offset=24
    i32.const 0
    local.get 2
    i32.const 24
    i32.add
    i32.store offset=65816
    local.get 2
    local.get 0
    i32.store
    local.get 2
    i32.const 0
    i32.load8_u offset=65652
    local.get 5
    call $runtime.hash32
    local.set 5
    local.get 3
    i32.const 0
    i32.load offset=65644
    local.tee 6
    i32.store
    local.get 6
    i32.const 0
    i32.load8_u offset=65653
    i32.const 0
    i32.load8_u offset=65652
    i32.add
    i32.const 3
    i32.shl
    i32.const 12
    i32.add
    local.get 5
    i32.const -1
    i32.const -1
    i32.const 0
    i32.load8_u offset=65654
    local.tee 3
    i32.shl
    i32.const -1
    i32.xor
    local.get 3
    i32.const 31
    i32.gt_u
    select
    i32.and
    i32.mul
    i32.add
    local.set 3
    local.get 5
    i32.const 24
    i32.shr_u
    local.tee 5
    i32.const 1
    local.get 5
    select
    local.set 7
    local.get 2
    i32.const 40
    i32.add
    local.set 8
    local.get 2
    i32.const 44
    i32.add
    local.set 9
    local.get 2
    i32.const 56
    i32.add
    local.set 10
    local.get 2
    i32.const 52
    i32.add
    local.set 11
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            loop  ;; label = @5
              local.get 8
              local.get 3
              i32.store
              local.get 9
              local.get 3
              i32.store
              local.get 2
              i32.const 24
              i32.add
              i32.const 12
              i32.add
              local.get 3
              i32.store
              local.get 3
              i32.eqz
              br_if 1 (;@4;)
              i32.const 0
              local.set 5
              loop  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    local.get 5
                    i32.const 8
                    i32.eq
                    br_if 0 (;@8;)
                    local.get 3
                    local.get 5
                    i32.add
                    i32.load8_u
                    local.get 7
                    i32.ne
                    br_if 1 (;@7;)
                    i32.const 0
                    i32.load8_u offset=65653
                    local.set 12
                    i32.const 0
                    i32.load8_u offset=65652
                    local.set 13
                    local.get 2
                    i32.const 24
                    i32.add
                    i32.const 24
                    i32.add
                    i32.const 0
                    i32.load offset=65656
                    local.tee 14
                    i32.store
                    local.get 11
                    i32.const 0
                    i32.load offset=65660
                    local.tee 6
                    i32.store
                    local.get 6
                    i32.eqz
                    br_if 6 (;@2;)
                    local.get 2
                    local.get 5
                    local.get 13
                    i32.mul
                    local.get 3
                    i32.add
                    i32.const 12
                    i32.add
                    i32.const 0
                    i32.load8_u offset=65652
                    local.get 14
                    local.get 6
                    call_indirect (type 0)
                    i32.const 1
                    i32.and
                    i32.eqz
                    br_if 1 (;@7;)
                    local.get 2
                    i32.const 8
                    i32.add
                    local.get 13
                    i32.const 3
                    i32.shl
                    local.get 5
                    local.get 12
                    i32.mul
                    i32.add
                    local.get 3
                    i32.add
                    i32.const 12
                    i32.add
                    i32.const 0
                    i32.load8_u offset=65653
                    call $memcpy
                    drop
                    br 5 (;@3;)
                  end
                  local.get 10
                  local.get 3
                  i32.load offset=8
                  local.tee 3
                  i32.store
                  br 2 (;@5;)
                end
                local.get 5
                i32.const 1
                i32.add
                local.set 5
                br 0 (;@6;)
              end
            end
          end
          local.get 2
          i32.const 8
          i32.add
          i32.const 0
          i32.const 0
          i32.load8_u offset=65653
          call $memset
          drop
        end
        local.get 2
        i32.const 104
        i32.add
        local.get 2
        i32.load offset=8
        local.tee 10
        i32.store
        local.get 2
        i32.const 100
        i32.add
        local.get 10
        i32.store
        local.get 2
        i32.const 88
        i32.add
        local.get 10
        i32.store
        local.get 2
        i32.const 84
        i32.add
        local.get 10
        i32.store
        local.get 2
        i32.load offset=12
        local.set 9
        local.get 2
        local.get 0
        i32.store offset=8
        local.get 2
        i32.const 8
        i32.add
        i32.const 0
        i32.load8_u offset=65652
        local.get 5
        call $runtime.hash32
        local.set 5
        local.get 2
        i32.const 60
        i32.add
        i32.const 0
        i32.load offset=65644
        local.tee 3
        i32.store
        local.get 3
        i32.const 0
        i32.load8_u offset=65653
        i32.const 0
        i32.load8_u offset=65652
        i32.add
        i32.const 3
        i32.shl
        i32.const 12
        i32.add
        local.get 5
        i32.const -1
        i32.const -1
        i32.const 0
        i32.load8_u offset=65654
        local.tee 6
        i32.shl
        i32.const -1
        i32.xor
        local.get 6
        i32.const 31
        i32.gt_u
        select
        i32.and
        i32.mul
        i32.add
        local.set 3
        local.get 5
        i32.const 24
        i32.shr_u
        local.tee 5
        i32.const 1
        local.get 5
        select
        local.set 7
        local.get 2
        i32.const 64
        i32.add
        local.set 0
        local.get 2
        i32.const 68
        i32.add
        local.set 15
        local.get 2
        i32.const 80
        i32.add
        local.set 16
        local.get 2
        i32.const 72
        i32.add
        local.set 12
        local.get 2
        i32.const 76
        i32.add
        local.set 8
        loop  ;; label = @3
          local.get 0
          local.get 3
          i32.store
          local.get 15
          local.get 3
          i32.store
          local.get 3
          i32.eqz
          br_if 2 (;@1;)
          i32.const 0
          local.set 5
          block  ;; label = @4
            loop  ;; label = @5
              local.get 5
              i32.const 8
              i32.eq
              br_if 1 (;@4;)
              block  ;; label = @6
                local.get 3
                local.get 5
                i32.add
                local.tee 13
                i32.load8_u
                local.get 7
                i32.ne
                br_if 0 (;@6;)
                i32.const 0
                i32.load8_u offset=65652
                local.set 14
                local.get 12
                i32.const 0
                i32.load offset=65656
                local.tee 11
                i32.store
                local.get 8
                i32.const 0
                i32.load offset=65660
                local.tee 6
                i32.store
                local.get 6
                i32.eqz
                br_if 4 (;@2;)
                local.get 2
                i32.const 8
                i32.add
                local.get 5
                local.get 14
                i32.mul
                local.get 3
                i32.add
                i32.const 12
                i32.add
                i32.const 0
                i32.load8_u offset=65652
                local.get 11
                local.get 6
                call_indirect (type 0)
                i32.const 1
                i32.and
                i32.eqz
                br_if 0 (;@6;)
                local.get 13
                i32.const 0
                i32.store8
                i32.const 0
                i32.const 0
                i32.load offset=65648
                i32.const -1
                i32.add
                i32.store offset=65648
                br 5 (;@1;)
              end
              local.get 5
              i32.const 1
              i32.add
              local.set 5
              br 0 (;@5;)
            end
          end
          local.get 16
          local.get 3
          i32.load offset=8
          local.tee 3
          i32.store
          br 0 (;@3;)
        end
      end
      call $runtime.nilPanic
      unreachable
    end
    local.get 2
    i32.const 108
    i32.add
    local.get 2
    i32.store
    local.get 2
    i32.const 92
    i32.add
    local.get 2
    i32.const 8
    i32.add
    i32.store
    local.get 2
    i32.const 0
    i32.store offset=16
    local.get 2
    i64.const 0
    i64.store offset=8
    local.get 2
    i32.const 136
    i32.add
    local.get 9
    call $runtime.alloc
    local.tee 5
    i32.store
    local.get 2
    i32.const 140
    i32.add
    local.get 5
    i32.store
    local.get 2
    i32.const 120
    i32.add
    local.get 5
    i32.store
    local.get 2
    i32.const 116
    i32.add
    local.get 5
    i32.store
    local.get 2
    i32.const 112
    i32.add
    local.get 5
    i32.store
    local.get 2
    i32.const 96
    i32.add
    local.get 5
    i32.store
    local.get 5
    local.get 10
    local.get 9
    call $memcpy
    local.set 3
    local.get 2
    i32.const 128
    i32.add
    local.get 2
    i32.const 8
    i32.add
    i32.store
    local.get 2
    i32.const 124
    i32.add
    local.get 2
    i32.store
    local.get 2
    i32.const 0
    i32.store offset=16
    local.get 2
    i64.const 0
    i64.store offset=8
    local.get 2
    i64.const 0
    i64.store
    local.get 2
    i32.const 152
    i32.add
    local.get 9
    call $runtime.alloc
    local.tee 5
    i32.store
    local.get 2
    i32.const 156
    i32.add
    local.get 5
    i32.store
    local.get 2
    i32.const 148
    i32.add
    local.get 5
    i32.store
    local.get 2
    i32.const 144
    i32.add
    local.get 5
    i32.store
    local.get 2
    i32.const 132
    i32.add
    local.get 5
    i32.store
    local.get 5
    local.get 3
    local.get 9
    call $memcpy
    local.set 5
    block  ;; label = @1
      local.get 9
      i32.eqz
      br_if 0 (;@1;)
      i32.const 0
      local.get 4
      i32.store offset=65816
      local.get 5
      local.get 9
      call $github.com/temphia/temphia/core/backend/server/engine/executors/wasm1/sdk._log
      local.get 2
      i32.const 160
      i32.add
      global.set $__stack_pointer
      return
    end
    call $runtime.lookupPanic
    unreachable)
  (func $strlen (type 5) (param i32) (result i32)
    (local i32 i32)
    local.get 0
    local.set 1
    block  ;; label = @1
      block  ;; label = @2
        local.get 0
        i32.const 3
        i32.and
        i32.eqz
        br_if 0 (;@2;)
        local.get 0
        local.set 1
        local.get 0
        i32.load8_u
        i32.eqz
        br_if 1 (;@1;)
        local.get 0
        i32.const 1
        i32.add
        local.tee 1
        i32.const 3
        i32.and
        i32.eqz
        br_if 0 (;@2;)
        local.get 1
        i32.load8_u
        i32.eqz
        br_if 1 (;@1;)
        local.get 0
        i32.const 2
        i32.add
        local.tee 1
        i32.const 3
        i32.and
        i32.eqz
        br_if 0 (;@2;)
        local.get 1
        i32.load8_u
        i32.eqz
        br_if 1 (;@1;)
        local.get 0
        i32.const 3
        i32.add
        local.tee 1
        i32.const 3
        i32.and
        i32.eqz
        br_if 0 (;@2;)
        local.get 1
        i32.load8_u
        i32.eqz
        br_if 1 (;@1;)
        local.get 0
        i32.const 4
        i32.add
        local.set 1
      end
      local.get 1
      i32.const -4
      i32.add
      local.set 1
      loop  ;; label = @2
        local.get 1
        i32.const 4
        i32.add
        local.tee 1
        i32.load
        local.tee 2
        i32.const -1
        i32.xor
        local.get 2
        i32.const -16843009
        i32.add
        i32.and
        i32.const -2139062144
        i32.and
        i32.eqz
        br_if 0 (;@2;)
      end
      local.get 2
      i32.const 255
      i32.and
      i32.eqz
      br_if 0 (;@1;)
      loop  ;; label = @2
        local.get 1
        i32.const 1
        i32.add
        local.tee 1
        i32.load8_u
        br_if 0 (;@2;)
      end
    end
    local.get 1
    local.get 0
    i32.sub)
  (func $memcpy (type 7) (param i32 i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i32)
    block  ;; label = @1
      block  ;; label = @2
        local.get 1
        i32.const 3
        i32.and
        i32.eqz
        br_if 0 (;@2;)
        local.get 2
        i32.eqz
        br_if 0 (;@2;)
        local.get 0
        local.get 1
        i32.load8_u
        i32.store8
        local.get 2
        i32.const -1
        i32.add
        local.set 3
        local.get 0
        i32.const 1
        i32.add
        local.set 4
        local.get 1
        i32.const 1
        i32.add
        local.tee 5
        i32.const 3
        i32.and
        i32.eqz
        br_if 1 (;@1;)
        local.get 3
        i32.eqz
        br_if 1 (;@1;)
        local.get 0
        local.get 1
        i32.load8_u offset=1
        i32.store8 offset=1
        local.get 2
        i32.const -2
        i32.add
        local.set 3
        local.get 0
        i32.const 2
        i32.add
        local.set 4
        local.get 1
        i32.const 2
        i32.add
        local.tee 5
        i32.const 3
        i32.and
        i32.eqz
        br_if 1 (;@1;)
        local.get 3
        i32.eqz
        br_if 1 (;@1;)
        local.get 0
        local.get 1
        i32.load8_u offset=2
        i32.store8 offset=2
        local.get 2
        i32.const -3
        i32.add
        local.set 3
        local.get 0
        i32.const 3
        i32.add
        local.set 4
        local.get 1
        i32.const 3
        i32.add
        local.tee 5
        i32.const 3
        i32.and
        i32.eqz
        br_if 1 (;@1;)
        local.get 3
        i32.eqz
        br_if 1 (;@1;)
        local.get 0
        local.get 1
        i32.load8_u offset=3
        i32.store8 offset=3
        local.get 2
        i32.const -4
        i32.add
        local.set 3
        local.get 0
        i32.const 4
        i32.add
        local.set 4
        local.get 1
        i32.const 4
        i32.add
        local.set 5
        br 1 (;@1;)
      end
      local.get 2
      local.set 3
      local.get 0
      local.set 4
      local.get 1
      local.set 5
    end
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 4
          i32.const 3
          i32.and
          local.tee 1
          br_if 0 (;@3;)
          block  ;; label = @4
            block  ;; label = @5
              local.get 3
              i32.const 16
              i32.lt_u
              br_if 0 (;@5;)
              block  ;; label = @6
                local.get 3
                i32.const -16
                i32.add
                local.tee 1
                i32.const 16
                i32.and
                br_if 0 (;@6;)
                local.get 4
                local.get 5
                i64.load align=4
                i64.store align=4
                local.get 4
                local.get 5
                i64.load offset=8 align=4
                i64.store offset=8 align=4
                local.get 4
                i32.const 16
                i32.add
                local.set 4
                local.get 5
                i32.const 16
                i32.add
                local.set 5
                local.get 1
                local.set 3
              end
              local.get 1
              i32.const 16
              i32.lt_u
              br_if 1 (;@4;)
              loop  ;; label = @6
                local.get 4
                local.get 5
                i64.load align=4
                i64.store align=4
                local.get 4
                i32.const 8
                i32.add
                local.get 5
                i32.const 8
                i32.add
                i64.load align=4
                i64.store align=4
                local.get 4
                i32.const 16
                i32.add
                local.get 5
                i32.const 16
                i32.add
                i64.load align=4
                i64.store align=4
                local.get 4
                i32.const 24
                i32.add
                local.get 5
                i32.const 24
                i32.add
                i64.load align=4
                i64.store align=4
                local.get 4
                i32.const 32
                i32.add
                local.set 4
                local.get 5
                i32.const 32
                i32.add
                local.set 5
                local.get 3
                i32.const -32
                i32.add
                local.tee 3
                i32.const 15
                i32.gt_u
                br_if 0 (;@6;)
              end
            end
            local.get 3
            local.set 1
          end
          block  ;; label = @4
            local.get 1
            i32.const 8
            i32.and
            i32.eqz
            br_if 0 (;@4;)
            local.get 4
            local.get 5
            i64.load align=4
            i64.store align=4
            local.get 5
            i32.const 8
            i32.add
            local.set 5
            local.get 4
            i32.const 8
            i32.add
            local.set 4
          end
          block  ;; label = @4
            local.get 1
            i32.const 4
            i32.and
            i32.eqz
            br_if 0 (;@4;)
            local.get 4
            local.get 5
            i32.load
            i32.store
            local.get 5
            i32.const 4
            i32.add
            local.set 5
            local.get 4
            i32.const 4
            i32.add
            local.set 4
          end
          block  ;; label = @4
            local.get 1
            i32.const 2
            i32.and
            i32.eqz
            br_if 0 (;@4;)
            local.get 4
            local.get 5
            i32.load16_u align=1
            i32.store16 align=1
            local.get 4
            i32.const 2
            i32.add
            local.set 4
            local.get 5
            i32.const 2
            i32.add
            local.set 5
          end
          local.get 1
          i32.const 1
          i32.and
          br_if 1 (;@2;)
          br 2 (;@1;)
        end
        block  ;; label = @3
          local.get 3
          i32.const 32
          i32.lt_u
          br_if 0 (;@3;)
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                local.get 1
                i32.const -1
                i32.add
                br_table 0 (;@6;) 1 (;@5;) 2 (;@4;) 3 (;@3;)
              end
              local.get 4
              local.get 5
              i32.load
              local.tee 6
              i32.store8
              local.get 4
              local.get 6
              i32.const 16
              i32.shr_u
              i32.store8 offset=2
              local.get 4
              local.get 6
              i32.const 8
              i32.shr_u
              i32.store8 offset=1
              local.get 3
              i32.const -3
              i32.add
              local.set 3
              local.get 4
              i32.const 3
              i32.add
              local.set 7
              i32.const 0
              local.set 1
              loop  ;; label = @6
                local.get 7
                local.get 1
                i32.add
                local.tee 4
                local.get 5
                local.get 1
                i32.add
                local.tee 2
                i32.const 4
                i32.add
                i32.load
                local.tee 8
                i32.const 8
                i32.shl
                local.get 6
                i32.const 24
                i32.shr_u
                i32.or
                i32.store
                local.get 4
                i32.const 4
                i32.add
                local.get 2
                i32.const 8
                i32.add
                i32.load
                local.tee 6
                i32.const 8
                i32.shl
                local.get 8
                i32.const 24
                i32.shr_u
                i32.or
                i32.store
                local.get 4
                i32.const 8
                i32.add
                local.get 2
                i32.const 12
                i32.add
                i32.load
                local.tee 8
                i32.const 8
                i32.shl
                local.get 6
                i32.const 24
                i32.shr_u
                i32.or
                i32.store
                local.get 4
                i32.const 12
                i32.add
                local.get 2
                i32.const 16
                i32.add
                i32.load
                local.tee 6
                i32.const 8
                i32.shl
                local.get 8
                i32.const 24
                i32.shr_u
                i32.or
                i32.store
                local.get 1
                i32.const 16
                i32.add
                local.set 1
                local.get 3
                i32.const -16
                i32.add
                local.tee 3
                i32.const 16
                i32.gt_u
                br_if 0 (;@6;)
              end
              local.get 7
              local.get 1
              i32.add
              local.set 4
              local.get 5
              local.get 1
              i32.add
              i32.const 3
              i32.add
              local.set 5
              br 2 (;@3;)
            end
            local.get 4
            local.get 5
            i32.load
            local.tee 6
            i32.store16 align=1
            local.get 3
            i32.const -2
            i32.add
            local.set 3
            local.get 4
            i32.const 2
            i32.add
            local.set 7
            i32.const 0
            local.set 1
            loop  ;; label = @5
              local.get 7
              local.get 1
              i32.add
              local.tee 4
              local.get 5
              local.get 1
              i32.add
              local.tee 2
              i32.const 4
              i32.add
              i32.load
              local.tee 8
              i32.const 16
              i32.shl
              local.get 6
              i32.const 16
              i32.shr_u
              i32.or
              i32.store
              local.get 4
              i32.const 4
              i32.add
              local.get 2
              i32.const 8
              i32.add
              i32.load
              local.tee 6
              i32.const 16
              i32.shl
              local.get 8
              i32.const 16
              i32.shr_u
              i32.or
              i32.store
              local.get 4
              i32.const 8
              i32.add
              local.get 2
              i32.const 12
              i32.add
              i32.load
              local.tee 8
              i32.const 16
              i32.shl
              local.get 6
              i32.const 16
              i32.shr_u
              i32.or
              i32.store
              local.get 4
              i32.const 12
              i32.add
              local.get 2
              i32.const 16
              i32.add
              i32.load
              local.tee 6
              i32.const 16
              i32.shl
              local.get 8
              i32.const 16
              i32.shr_u
              i32.or
              i32.store
              local.get 1
              i32.const 16
              i32.add
              local.set 1
              local.get 3
              i32.const -16
              i32.add
              local.tee 3
              i32.const 17
              i32.gt_u
              br_if 0 (;@5;)
            end
            local.get 7
            local.get 1
            i32.add
            local.set 4
            local.get 5
            local.get 1
            i32.add
            i32.const 2
            i32.add
            local.set 5
            br 1 (;@3;)
          end
          local.get 4
          local.get 5
          i32.load
          local.tee 6
          i32.store8
          local.get 3
          i32.const -1
          i32.add
          local.set 3
          local.get 4
          i32.const 1
          i32.add
          local.set 7
          i32.const 0
          local.set 1
          loop  ;; label = @4
            local.get 7
            local.get 1
            i32.add
            local.tee 4
            local.get 5
            local.get 1
            i32.add
            local.tee 2
            i32.const 4
            i32.add
            i32.load
            local.tee 8
            i32.const 24
            i32.shl
            local.get 6
            i32.const 8
            i32.shr_u
            i32.or
            i32.store
            local.get 4
            i32.const 4
            i32.add
            local.get 2
            i32.const 8
            i32.add
            i32.load
            local.tee 6
            i32.const 24
            i32.shl
            local.get 8
            i32.const 8
            i32.shr_u
            i32.or
            i32.store
            local.get 4
            i32.const 8
            i32.add
            local.get 2
            i32.const 12
            i32.add
            i32.load
            local.tee 8
            i32.const 24
            i32.shl
            local.get 6
            i32.const 8
            i32.shr_u
            i32.or
            i32.store
            local.get 4
            i32.const 12
            i32.add
            local.get 2
            i32.const 16
            i32.add
            i32.load
            local.tee 6
            i32.const 24
            i32.shl
            local.get 8
            i32.const 8
            i32.shr_u
            i32.or
            i32.store
            local.get 1
            i32.const 16
            i32.add
            local.set 1
            local.get 3
            i32.const -16
            i32.add
            local.tee 3
            i32.const 18
            i32.gt_u
            br_if 0 (;@4;)
          end
          local.get 7
          local.get 1
          i32.add
          local.set 4
          local.get 5
          local.get 1
          i32.add
          i32.const 1
          i32.add
          local.set 5
        end
        block  ;; label = @3
          local.get 3
          i32.const 16
          i32.and
          i32.eqz
          br_if 0 (;@3;)
          local.get 4
          local.get 5
          i32.load8_u
          i32.store8
          local.get 4
          local.get 5
          i32.load offset=1 align=1
          i32.store offset=1 align=1
          local.get 4
          local.get 5
          i64.load offset=5 align=1
          i64.store offset=5 align=1
          local.get 4
          local.get 5
          i32.load16_u offset=13 align=1
          i32.store16 offset=13 align=1
          local.get 4
          local.get 5
          i32.load8_u offset=15
          i32.store8 offset=15
          local.get 4
          i32.const 16
          i32.add
          local.set 4
          local.get 5
          i32.const 16
          i32.add
          local.set 5
        end
        block  ;; label = @3
          local.get 3
          i32.const 8
          i32.and
          i32.eqz
          br_if 0 (;@3;)
          local.get 4
          local.get 5
          i64.load align=1
          i64.store align=1
          local.get 4
          i32.const 8
          i32.add
          local.set 4
          local.get 5
          i32.const 8
          i32.add
          local.set 5
        end
        block  ;; label = @3
          local.get 3
          i32.const 4
          i32.and
          i32.eqz
          br_if 0 (;@3;)
          local.get 4
          local.get 5
          i32.load align=1
          i32.store align=1
          local.get 4
          i32.const 4
          i32.add
          local.set 4
          local.get 5
          i32.const 4
          i32.add
          local.set 5
        end
        block  ;; label = @3
          local.get 3
          i32.const 2
          i32.and
          i32.eqz
          br_if 0 (;@3;)
          local.get 4
          local.get 5
          i32.load16_u align=1
          i32.store16 align=1
          local.get 4
          i32.const 2
          i32.add
          local.set 4
          local.get 5
          i32.const 2
          i32.add
          local.set 5
        end
        local.get 3
        i32.const 1
        i32.and
        i32.eqz
        br_if 1 (;@1;)
      end
      local.get 4
      local.get 5
      i32.load8_u
      i32.store8
    end
    local.get 0)
  (func $memset (type 7) (param i32 i32 i32) (result i32)
    (local i32 i32 i32 i64)
    block  ;; label = @1
      local.get 2
      i32.eqz
      br_if 0 (;@1;)
      local.get 0
      local.get 1
      i32.store8
      local.get 2
      local.get 0
      i32.add
      local.tee 3
      i32.const -1
      i32.add
      local.get 1
      i32.store8
      local.get 2
      i32.const 3
      i32.lt_u
      br_if 0 (;@1;)
      local.get 0
      local.get 1
      i32.store8 offset=2
      local.get 0
      local.get 1
      i32.store8 offset=1
      local.get 3
      i32.const -3
      i32.add
      local.get 1
      i32.store8
      local.get 3
      i32.const -2
      i32.add
      local.get 1
      i32.store8
      local.get 2
      i32.const 7
      i32.lt_u
      br_if 0 (;@1;)
      local.get 0
      local.get 1
      i32.store8 offset=3
      local.get 3
      i32.const -4
      i32.add
      local.get 1
      i32.store8
      local.get 2
      i32.const 9
      i32.lt_u
      br_if 0 (;@1;)
      local.get 0
      i32.const 0
      local.get 0
      i32.sub
      i32.const 3
      i32.and
      local.tee 4
      i32.add
      local.tee 3
      local.get 1
      i32.const 255
      i32.and
      i32.const 16843009
      i32.mul
      local.tee 1
      i32.store
      local.get 3
      local.get 2
      local.get 4
      i32.sub
      i32.const -4
      i32.and
      local.tee 4
      i32.add
      local.tee 2
      i32.const -4
      i32.add
      local.get 1
      i32.store
      local.get 4
      i32.const 9
      i32.lt_u
      br_if 0 (;@1;)
      local.get 3
      local.get 1
      i32.store offset=8
      local.get 3
      local.get 1
      i32.store offset=4
      local.get 2
      i32.const -8
      i32.add
      local.get 1
      i32.store
      local.get 2
      i32.const -12
      i32.add
      local.get 1
      i32.store
      local.get 4
      i32.const 25
      i32.lt_u
      br_if 0 (;@1;)
      local.get 3
      local.get 1
      i32.store offset=24
      local.get 3
      local.get 1
      i32.store offset=20
      local.get 3
      local.get 1
      i32.store offset=16
      local.get 3
      local.get 1
      i32.store offset=12
      local.get 2
      i32.const -16
      i32.add
      local.get 1
      i32.store
      local.get 2
      i32.const -20
      i32.add
      local.get 1
      i32.store
      local.get 2
      i32.const -24
      i32.add
      local.get 1
      i32.store
      local.get 2
      i32.const -28
      i32.add
      local.get 1
      i32.store
      local.get 4
      local.get 3
      i32.const 4
      i32.and
      i32.const 24
      i32.or
      local.tee 5
      i32.sub
      local.tee 2
      i32.const 32
      i32.lt_u
      br_if 0 (;@1;)
      local.get 1
      i64.extend_i32_u
      i64.const 4294967297
      i64.mul
      local.set 6
      local.get 3
      local.get 5
      i32.add
      local.set 1
      loop  ;; label = @2
        local.get 1
        local.get 6
        i64.store
        local.get 1
        i32.const 24
        i32.add
        local.get 6
        i64.store
        local.get 1
        i32.const 16
        i32.add
        local.get 6
        i64.store
        local.get 1
        i32.const 8
        i32.add
        local.get 6
        i64.store
        local.get 1
        i32.const 32
        i32.add
        local.set 1
        local.get 2
        i32.const -32
        i32.add
        local.tee 2
        i32.const 31
        i32.gt_u
        br_if 0 (;@2;)
      end
    end
    local.get 0)
  (table (;0;) 3 3 funcref)
  (memory (;0;) 2)
  (global $__stack_pointer (mut i32) (i32.const 65536))
  (export "memory" (memory 0))
  (export "malloc" (func $malloc))
  (export "free" (func $free))
  (export "calloc" (func $calloc))
  (export "realloc" (func $realloc))
  (export "_start" (func $_start))
  (export "action_main" (func $action_main))
  (elem (;0;) (i32.const 1) func $runtime.memequal $runtime.hash32)
  (data $.rodata (i32.const 65536) "out of memorypanic: runtime error: nil pointer dereferenceindex out of rangeslice out of range")
  (data $.data (i32.const 65632) "\10\02\01\00\90\00\01\00\00\00\00\00|\01\01\00\00\00\00\00\04\0c\00\00\00\00\00\00\01\00\00\00\00\00\00\00\02\00\00\00"))
